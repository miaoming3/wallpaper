package server

import (
	"github.com/miaoming3/wallpaper/response"
)

type BaseServiceInterface interface {
	IndexServer(data interface{}) *response.ApiResponse
	UpdateServer(data interface{}) *response.ApiResponse
	CreateServer(data interface{}) *response.ApiResponse
	DeleteServer(data interface{}) *response.ApiResponse
}
