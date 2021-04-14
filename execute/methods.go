package execute

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/lil5/go-site-bk/config"
)

func MakedirSiteCache(c *config.Config, siteIndex int) error {
	folderPath := fmt.Sprintf(".cache/%s", c.Site[siteIndex].Name)
	info, err := os.Stat(folderPath)

	if info != nil && !info.IsDir() {
		return fmt.Errorf("Folder does not exist. %s", folderPath)
	}

	if os.IsNotExist(err) {
		runExecf(
			"mkdir -p %s",
			folderPath,
		)
	}

	return nil
}

func BackupDatabase(c *config.Config, siteIndex int) error {
	return runExecf(
		"ssh %s@%s -p %d \"mysqldump -u %s -p%s --databases %s\" > .cache/%s/db.sql",
		c.Site[siteIndex].Ssh.User,
		c.Site[siteIndex].Ssh.Host,
		c.Site[siteIndex].Ssh.Port,
		c.Site[siteIndex].Mysql.User,
		c.Site[siteIndex].Mysql.Pass,
		strings.Join(c.Site[siteIndex].Mysql.Databases, " "),
		c.Site[siteIndex].Name,
	)
}

func BackupFiles(c *config.Config, siteIndex int) error {
	if c.Site[siteIndex].Rsync.Www == "" {
		log.Printf("Site: %s wil not use rsync\n\n", c.Site[siteIndex].Name)
		return nil
	}

	var excludeStr string
	for _, ex := range c.Site[siteIndex].Rsync.Exclude {
		excludeStr += fmt.Sprintf(" --exclude='%s'", ex)
	}

	return runExecf(
		"rsync --progress -az --delete -e \"ssh -p %d\"%s %s@%s:%s .cache/%s/",
		c.Site[siteIndex].Ssh.Port,
		excludeStr,
		c.Site[siteIndex].Ssh.User,
		c.Site[siteIndex].Ssh.Host,
		strings.TrimSuffix(c.Site[siteIndex].Rsync.Www, "/"),
		c.Site[siteIndex].Name,
	)
}

func TarEncrypt(c *config.Config, siteIndex int) error {
	return runExecf(
		"cd .cache/%s;TIMESTAMP=`date \"+%%Y-%%m-%%d-%%H%%M%%S\"`; tar -czf - * | gpg -q -c --batch --passphrase %s > ../../archive-%s-$TIMESTAMP.tar.gz.gpg; cd ..",
		c.Site[siteIndex].Name,
		c.Site[siteIndex].Encrypt,
		c.Site[siteIndex].Name,
	)
}
