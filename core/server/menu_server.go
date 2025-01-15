package server

import (
	"github.com/gin-gonic/gin"
	response2 "github.com/miaoming3/wallpaper/http/response"
	"github.com/miaoming3/wallpaper/response"
)

type MenuServer struct {
}

func NewMenuServe() BaseServiceInterface {
	return &MenuServer{}
}

func (ms *MenuServer) IndexServer(c *gin.Context, di interface{}) *response.ApiResponse {

	return response2.ApiSuccess(nil)
}

func (ms *MenuServer) UpdateServer(c *gin.Context, di interface{}) *response.ApiResponse {
	return response2.ApiSuccess(nil)
}

func (ms *MenuServer) CreateServer(c *gin.Context, di interface{}) *response.ApiResponse {
	return response2.ApiSuccess(nil)
}

func (ms *MenuServer) DeleteServer(di interface{}) *response.ApiResponse {
	return response2.ApiSuccess(nil)
}
