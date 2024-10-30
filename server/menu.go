package server

import (
	"github.com/gin-gonic/gin"
	"github.com/miaoming3/wallpaper/response"
)

type MenuService interface {
	IndexServer(c *gin.Context) *response.ApiResponse
	UpdateServer(c *gin.Context) *response.ApiResponse
	CreateServer(c *gin.Context) *response.ApiResponse
	DeleteServer(c *gin.Context) *response.ApiResponse
}
type MenuServer struct {
}

func NewMenuServe() MenuService {
	return &MenuServer{}
}

func (ms *MenuServer) IndexServer(c *gin.Context) *response.ApiResponse {

	return response.ApiSuccess(nil)
}

func (ms *MenuServer) UpdateServer(c *gin.Context) *response.ApiResponse {
	return response.ApiSuccess(nil)
}

func (ms *MenuServer) CreateServer(c *gin.Context) *response.ApiResponse {
	return response.ApiSuccess(nil)
}

func (ms *MenuServer) DeleteServer(c *gin.Context) *response.ApiResponse {
	return response.ApiSuccess(nil)
}
