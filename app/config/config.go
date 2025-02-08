package config

type Config struct {
	Model           string `json:"model" yaml:"model" mapstructure:"model"`
	SysPort         int    `json:"sys_port" yaml:"sys_port" mapstructure:"sys_port"`
	*DataBaseConfig `yaml:"data_base" json:"data_base" mapstructure:"data_base"`
	*UploadConfig   `yaml:"upload_config" json:"upload_config" mapstructure:"upload_config"`
	*CaptchaConfig  `yaml:"captcha_config" json:"captcha_config" mapstructure:"captcha_config"`
	*LoggerConfig   `yaml:"logger_config" mapstructure:"logger_config"`
	RedisConfig     *RedisConfig `yaml:"redis_config" mapstructure:"redis_config"`
	AesKey          string       `yaml:"aes_key" mapstructure:"aes_key"`
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

type CaptchaConfig struct {
	CaptchaType     string   `json:"captcha_type" yaml:"captcha_type" mapstructure:"captcha_type"`
	Width           int      `yaml:"width" json:"width" mapstructure:"width"`
	Height          int      `yaml:"height" json:"height" mapstructure:"height"`
	NoiseCount      int      `yaml:"noise_count" json:"noise_count" mapstructure:"noise_count"`
	ShowLineOptions int      `json:"show_line_options" yaml:"show_line_options" mapstructrue:"show_line_options"`
	Length          int      `json:"length" yaml:"length" mapstructrue:"length"`
	Source          string   `json:"source" yaml:"source" mapstructrue:"source"`
	BgColor         []uint8  `json:"bg_color" yaml:"bg_color" mapstructrue:"bg_color"`
	Fonts           []string `json:"fonts" yaml:"fonts" mapstructrue:"fonts"`
}

type LoggerConfig struct {
	Prefix       string `yaml:"prefix" mapstructure:"prefix"`                 // 前缀
	Director     string `yaml:"director" mapstructure:"director"`             // 日志存放路径
	ShowLine     bool   `yaml:"show_line" mapstructure:"show_line"`           // 是否显示行号
	Level        int8   `yaml:"level" mapstructure:"level"`                   // 是否显示行号
	LogInConsole bool   `yaml:"log_in_console" mapstructure:"log_in_console"` // 是否显示打印日志
}

type RedisConfig struct {
	Username string `yaml:"username" mapstructure:"username" `
	Password string `yaml:"password" mapstructure:"password" `
	Host     string `yaml:"host" mapstructure:"host"`
	Port     int    `yaml:"port" mapstructure:"port"`
	Db       int    `yaml:"db" mapstructure:"db"`
}
