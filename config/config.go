package config

type Config struct {
	AllowedExtensions []string `json:"allowed_extensions" yaml:"allowed_extensions" mapstructure:"allowed_extensions"`
	UploadFileSize    int64    `json:"upload_file_size" yaml:"upload_file_size" mapstructure:"upload_file_size"`
	UploadDir         string   `json:"upload_dir" yaml:"upload_dir" mapstructure:"upload_dir"`
	Template          string   `json:"template" yaml:"template" mapstructure:"template"`
	Dev               bool     `json:"dev" yaml:"dev" mapstructure:"dev"`
	SysPort           int      `json:"sys_port" yaml:"sys_port" mapstructure:"sys_port"`
	*DataBaseConfig   `yaml:"data_base" json:"data_base" mapstructure:"data_base"`
	*UploadConfig     `yaml:"upload_config" json:"upload_config" mapstructure:"upload_config"`
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
	Platform    string `json:"platform" yaml:"platform" mapstructure:"platform"`
	OssPlatform string `json:"oss_platform" yaml:"oss_platform" mapstructure:"oss_platform"`
	AccessKey   string `json:"access_key" yaml:"access_key" mapstructure:"access_key"`
	SecretKey   string `json:"secret_key" yaml:"secret_key" mapstructure:"secret_key"`
	Bucket      string `json:"bucket" yaml:"bucket" mapstructure:"bucket"`
	Region      string `json:"region" yaml:"region" mapstructure:"region"`
	Endpoint    string `json:"endpoint" yaml:"endpoint" mapstructure:"endpoint"`
	Dir         string `json:"dir" yaml:"dir" mapstructure:"dir"`
	Url         string `json:"url" yaml:"url" mapstructure:"url"`
	Method      string `json:"method" yaml:"method" mapstructure:"method"`
	Notice      string `json:"notice" yaml:"notice" mapstructure:"notice"`
}
