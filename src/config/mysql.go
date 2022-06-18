package config

type Mysql struct {
	Path     string `mapstructure:"path" json:"path,omitempty" yaml:"path"`
	Port     string `mapstructure:"port" json:"port,omitempty" yaml:"port"`
	Dbname   string `mapstructure:"db-name" json:"db-name,omitempty" yaml:"db-name"`
	Username string `mapstructure:"username" json:"username,omitempty" yaml:"username"`
	Password string `mapstructure:"password" json:"password,omitempty" yaml:"password"`
	Config   string `mapstructure:"config" json:"config" yaml:"config"`
}
