package config

type Config struct {
	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
}

func (mysql *Mysql) Dsn() string {
	return mysql.Username + ":" + mysql.Password + "@tcp(" + mysql.Path + ":" + mysql.Port + ")/" + mysql.Dbname + mysql.Config
}
