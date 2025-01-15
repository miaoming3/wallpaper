package server

import (
	"github.com/gin-gonic/gin"
	"github.com/miaoming3/wallpaper/app/response"
)

type BaseServiceInterface interface {
	IndexServer(c *gin.Context, data interface{}) *response.APi
	UpdateServer(c *gin.Context, data interface{}) *response.APi
	CreateServer(c *gin.Context, data interface{}) *response.APi
	DeleteServer(data interface{}) *response.APi
}
