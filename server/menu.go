package server

import (
	"github.com/miaoming3/wallpaper/response"
)

type MenuServer struct {
}

func NewMenuServe() BaseServiceInterface {
	return &MenuServer{}
}

func (ms *MenuServer) IndexServer(di interface{}) *response.ApiResponse {

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
