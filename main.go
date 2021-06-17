package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/lil5/go-site-bk/config"
	"github.com/lil5/go-site-bk/execute"
	"github.com/urfave/cli/v2"
)

func checkMissingProg() error {
	var missing []string
	var err error

	progs := []string{
		"rsync",
		"ssh",
		"tar",
	}

	for _, prog := range progs {
		cmd := exec.Command("which", prog)
		if cmd.Run() != nil {
			missing = append(missing, prog)
			break
		}
	}

	if len(missing) > 0 {
		err = fmt.Errorf("Error: [%s] does not exist", strings.Join(missing, ", "))
	}

	return err
}

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func run(filterByName []string) error {
	if runtime.GOOS == "windows" {
		return fmt.Errorf("For now this program only works on unix systems")
	}

	err := checkMissingProg()
	if err != nil {
		return err
	}

	c, err := config.GetConfig()
	if err != nil {
		return err
	}

	allowed, err := config.CheckConfig(c)
	if err != nil {
		return err
	}

	for siteIndex := range c.Site {
		if !allowed[siteIndex].Rsync && !allowed[siteIndex].Mysql {
			continue
		}

		if len(filterByName) > 0 && !contains(filterByName, c.Site[siteIndex].Name) {
			continue
		}

		fmt.Printf("Site: %s\n", c.Site[siteIndex].Name)

		err := execute.MakedirSiteCache(c, siteIndex)
		if err != nil {
			return err
		}

		if allowed[siteIndex].Mysql {
			fmt.Print("Mysql ...")
			err = execute.BackupDatabase(c, siteIndex)

			if err != nil {
				return err
			}
			fmt.Println(" Done.")
		}

		if allowed[siteIndex].Rsync {
			fmt.Print("Rsync ...")
			err = execute.BackupFiles(c, siteIndex)

			if err != nil {
				return err
			}
			fmt.Println(" Done.")
		}

		fmt.Print("Tar Encrypt ...")
		err = execute.TarEncrypt(c, siteIndex)
		if err != nil {
			return err
		}
		fmt.Print(" Done.\n\n")
	}

	return nil
}

func main() {
	app := &cli.App{
		Name:                 "Site Backup",
		Usage:                "Create a config.toml and run",
		ArgsUsage:            "A list of site names to filter what to sync",
		EnableBashCompletion: true,
		Authors: []*cli.Author{
			{
				Name:  "Lucian I. Last",
				Email: "li@last.nl",
			},
		},
		Action: func(c *cli.Context) error {
			argsArr := c.Args().Slice()

			err := run(argsArr)

			if err != nil {
				return cli.Exit(err, 2)
			}
			return nil
		},
	}

	app.Run(os.Args)
}
