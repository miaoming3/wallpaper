package server

import (
	"github.com/miaoming3/wallpaper/controller/dto"
	"github.com/miaoming3/wallpaper/response"
)

type ImageServer struct {
}

func NewImageServer() BaseServiceInterface {
	return &ImageServer{}
}

func (is *ImageServer) IndexServer(di interface{}) *response.ApiResponse {
	data, ok := di.(dto.ImageSearch)
	if !ok {
		return response.ApiError(response.ACCESSERROR, nil)
	}
	if data.CID != 0 {

	}
	return response.ApiSuccess(nil)
}
func (is *ImageServer) UpdateServer(di interface{}) *response.ApiResponse {
	return response.ApiSuccess(nil)
}
func (is *ImageServer) CreateServer(di interface{}) *response.ApiResponse {
	return response.ApiSuccess(nil)
}
func (is *ImageServer) DeleteServer(di interface{}) *response.ApiResponse {
	return response.ApiSuccess(nil)
}
