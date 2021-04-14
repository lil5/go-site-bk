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
		err = fmt.Errorf("Error: %s does not exist", strings.Join(missing, ", "))
	}

	return err
}

func run() error {
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

		fmt.Printf("Site: %s", c.Site[siteIndex].Name)

		if allowed[siteIndex].Mysql {
			err = execute.BackupDatabase(c, siteIndex)

			if err != nil {
				return err
			}
		}
		if allowed[siteIndex].Rsync {
			err = execute.BackupFiles(c, siteIndex)

			if err != nil {
				return err
			}
		}

		err = execute.TarEncrypt(c, siteIndex)
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	app := &cli.App{
		Name:                 "Site Backup",
		Usage:                "Create a config.toml and run",
		EnableBashCompletion: true,
		Authors: []*cli.Author{
			{
				Name:  "Lucian I. Last",
				Email: "li@last.nl",
			},
		},
		Action: func(c *cli.Context) error {
			err := run()

			if err != nil {
				return cli.Exit(err, 1)
			}
			return nil
		},
	}

	app.Run(os.Args)
}
