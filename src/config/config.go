package config

type Config struct {
	Mysql  Mysql  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Casbin Casbin `mapstructure:"casbin" json:"casbin" yaml:"casbin"`
	Jwt    Jwt    `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
}

func (mysql *Mysql) Dsn() string {
	return mysql.Username + ":" + mysql.Password + "@tcp(" + mysql.Path + ":" + mysql.Port + ")/" + mysql.Dbname + mysql.Config
}
