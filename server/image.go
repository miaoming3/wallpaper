package server

import (
	"github.com/gin-gonic/gin"
	"github.com/miaoming3/wallpaper/controller/dto"
	"github.com/miaoming3/wallpaper/dao"
	"github.com/miaoming3/wallpaper/response"
)

type ImageServer struct {
}

func NewImageServer() BaseServiceInterface {
	return &ImageServer{}
}

func (is *ImageServer) IndexServer(c *gin.Context, di interface{}) *response.ApiResponse {
	data, ok := di.(*dto.ImageSearch)
	if !ok {
		return response.ApiError(response.ACCESSERROR, nil)
	}
	condition := make(map[string]interface{})
	if data.CID != 0 {
		condition["cid"] = data.CID
	}
	if data.Username != "" {
		condition["username like ?"] = "%" + data.Username + "%"
	}
	if data.Name != "" {
		condition["name like ?"] = "%" + data.Name + "%"
	}
	if data.TagsID > 0 {
		condition["tags_id"] = data.TagsID
	}
	if data.IsRecommend > 0 {
		condition["is_recommend"] = data.IsRecommend
	}
	page := c.GetInt("page")
	pageSize := c.GetInt("pageSize")
	images, err := dao.NewImagesDao().FindByAll(condition, page, pageSize)
	if err != nil {
		return response.ApiError(response.ACCESSERROR, err)
	}
	total, err := dao.NewImagesDao().FindByTotal(condition)
	if err != nil {
		return response.ApiError(response.ACCESSERROR, err)
	}
	return response.ApiPageSuccess(images, total, page, pageSize, total/int64(pageSize) > int64(page))
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