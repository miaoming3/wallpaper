package server

import (
	"github.com/gin-gonic/gin"
	"github.com/miaoming3/wallpaper/app/response"
	response2 "github.com/miaoming3/wallpaper/app/response"
)

type MenuServer struct {
}

func NewMenuServe() BaseServiceInterface {
	return &MenuServer{}
}

func (ms *MenuServer) IndexServer(c *gin.Context, di interface{}) *response.APi {

	return response2.ApiSuccess(nil)
}

func (ms *MenuServer) UpdateServer(c *gin.Context, di interface{}) *response.APi {
	return response2.ApiSuccess(nil)
}

func (ms *MenuServer) CreateServer(c *gin.Context, di interface{}) *response.APi {
	return response2.ApiSuccess(nil)
}

func (ms *MenuServer) DeleteServer(di interface{}) *response.APi {
	return response2.ApiSuccess(nil)
}
