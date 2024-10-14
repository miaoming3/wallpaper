package config

type Config struct {
	Dev             bool `json:"dev" yaml:"dev" mapstructure:"dev"`
	SysPort         int  `json:"sys_port" yaml:"sys_port" mapstructure:"sys_port"`
	*DataBaseConfig `yaml:"data_base" json:"data_base" mapstructure:"data_base"`
}

type DataBaseConfig struct {
	Host     string `json:"host" yaml:"host" mapstructure:"host"`
	Password string `json:"password" yaml:"password" mapstructure:"password"`
	DbName   string `json:"db_name" yaml:"db_name" mapstructure:"db_name"`
	Username string `json:"username" yaml:"username" mapstructure:"password"`
	Charset  string `json:"charset" yaml:"charset" mapstructure:"charset"`
	Port     int    `json:"port" yaml:"port" mapstucture:"post"`
	Prefix   string `json:"prefix" yaml:"prefix" mapstucture:"prefix"`
}
