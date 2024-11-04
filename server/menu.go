package server

import (
	"github.com/gin-gonic/gin"
	"github.com/miaoming3/wallpaper/response"
)

type MenuServer struct {
}

func NewMenuServe() BaseServiceInterface {
	return &MenuServer{}
}

func (ms *MenuServer) IndexServer(c *gin.Context, di interface{}) *response.ApiResponse {

	return response.ApiSuccess(nil)
}

func (ms *MenuServer) UpdateServer(di interface{}) *response.ApiResponse {
	return response.ApiSuccess(nil)
}

func (ms *MenuServer) CreateServer(di interface{}) *response.ApiResponse {
	return response.ApiSuccess(nil)
}

func (ms *MenuServer) DeleteServer(di interface{}) *response.ApiResponse {
	return response.ApiSuccess(nil)
}
