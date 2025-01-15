package server

import (
	"github.com/gin-gonic/gin"
	"github.com/miaoming3/wallpaper/app/response"
)

type BaseServiceInterface interface {
	IndexServer(c *gin.Context, data interface{}) *response.ApiResponse
	UpdateServer(c *gin.Context, data interface{}) *response.ApiResponse
	CreateServer(c *gin.Context, data interface{}) *response.ApiResponse
	DeleteServer(data interface{}) *response.ApiResponse
}
