package server

import (
	"github.com/gin-gonic/gin"
	"github.com/miaoming3/wallpaper/response"
)

type BaseServiceInterface interface {
	IndexServer(c *gin.Context, data interface{}) *response.ApiResponse
	UpdateServer(data interface{}) *response.ApiResponse
	CreateServer(data interface{}) *response.ApiResponse
	DeleteServer(data interface{}) *response.ApiResponse
}
