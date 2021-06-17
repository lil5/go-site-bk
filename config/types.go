package config

type Config struct {
	Site []Site `toml:"site"`
}
type Site struct {
	Name    string
	Encrypt string
	Local   bool
	Ssh     Ssh
	Rsync   Rsync
	Mysql   Mysql
}
type Rsync struct {
	Www     string
	Exclude []string
}
type Ssh struct {
	Host string
	User string
	Port uint
	Pass string
}
type Mysql struct {
	Databases []string
	User      string
	Pass      string
}

type AllowedRuns struct {
	Rsync bool
	Mysql bool
}
