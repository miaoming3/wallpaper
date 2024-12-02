package filesystem

import (
	"context"
)

type UploadConfig struct {
	config *Config
}

type Config struct {
	Platform    string
	OssPlatform string
	AccessKey   string
	SecretKey   string
	Bucket      string
	Region      string
	Endpoint    string
	Dir         string
	Url         string
	Method      string
	Notice      string
}

type Interface interface {
	UploadFile(c *context.Context, data ...interface{}) interface{}
}
