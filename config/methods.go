package config

import (
	"fmt"
	"log"

	toml "github.com/pelletier/go-toml"
)

const (
	ErrorMissing = "Missing site[%d]%s from config.toml"
)

func GetConfig() (*Config, error) {
	tree, err := toml.LoadFile("config.toml")
	if err != nil {
		log.Printf("Fatal error config file: %s \n", err)
		return nil, err
	}

	var c Config
	err = tree.Unmarshal(&c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

// Check config for missing values and add default values where ommitted
func CheckConfig(c *Config) ([]AllowedRuns, error) {
	allowed := make([]AllowedRuns, len(c.Site))
	var err error
	for siteIndex := range c.Site {
		// check basic requirements
		if c.Site[siteIndex].Name == "" {
			err = fmt.Errorf(ErrorMissing, siteIndex, "name")
		}
		if c.Site[siteIndex].Ssh.Host == "" {
			err = fmt.Errorf(ErrorMissing, siteIndex, "ssh.host")
		}
		if c.Site[siteIndex].Ssh.User == "" {
			err = fmt.Errorf(ErrorMissing, siteIndex, "ssh.host")
		}

		if err != nil {
			return nil, err
		}

		okMysql, err := checkMysql(c, siteIndex)
		if err != nil {
			return nil, err
		}
		okRsync, err := checkRsync(c, siteIndex)
		if err != nil {
			return nil, err
		}

		// set defaults
		if c.Site[siteIndex].Ssh.Port == 0 {
			c.Site[siteIndex].Ssh.Port = 22
		}

		allowed[siteIndex] = AllowedRuns{
			Mysql: okMysql,
			Rsync: okRsync,
		}
	}

	return allowed, nil
}

func checkRsync(c *Config, siteIndex int) (bool, error) {
	emptyWww := c.Site[siteIndex].Rsync.Www == ""
	emptyExclude := len(c.Site[siteIndex].Rsync.Exclude) == 0
	if emptyWww && emptyExclude {
		return false, nil
	}
	if !emptyWww {
		return true, nil
	}

	err := fmt.Errorf(ErrorMissing, siteIndex, "rsync.www")
	return false, err
}

func checkMysql(c *Config, siteIndex int) (bool, error) {
	emptyDatabases := len(c.Site[siteIndex].Mysql.Databases) == 0
	emptyPass := c.Site[siteIndex].Mysql.Pass == ""
	emptyUser := c.Site[siteIndex].Mysql.User == ""

	if emptyDatabases && emptyPass && emptyUser {
		return false, nil
	}
	if !emptyPass && !emptyUser {
		return true, nil
	}

	err := fmt.Errorf(ErrorMissing, siteIndex, "mysql.[pass,user]")
	return false, err
}
