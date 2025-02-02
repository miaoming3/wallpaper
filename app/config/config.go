package config

type Config struct {
	Model           string `json:"model" yaml:"model" mapstructure:"model"`
	SysPort         int    `json:"sys_port" yaml:"sys_port" mapstructure:"sys_port"`
	*DataBaseConfig `yaml:"data_base" json:"data_base" mapstructure:"data_base"`
	*UploadConfig   `yaml:"upload_config" json:"upload_config" mapstructure:"upload_config"`
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

type UploadConfig struct {
	Dir               string   `json:"dir" yaml:"dir" mapstructure:"dir"`
	AllowedExtensions []string `json:"allowed_extensions" yaml:"allowed_extensions" mapstructure:"allowed_extensions"`
	UploadFileSize    int64    `json:"upload_file_size" yaml:"upload_file_size" mapstructure:"upload_file_size"`
}
