package server

import (
	"github.com/gin-gonic/gin"
	"github.com/miaoming3/wallpaper/controller/dto"
	"github.com/miaoming3/wallpaper/dao"
	"github.com/miaoming3/wallpaper/response"
	"github.com/miaoming3/wallpaper/response/dro"
	"github.com/miaoming3/wallpaper/utils"
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
	condition := dao.NewQueryOption()
	condition.AddPreload("Category", nil)
	condition.AddPreload("User", nil)
	condition.AddPreload("Tags", nil)
	if data.CID != 0 {
		condition.AddCondition("cid = ?", data.CID)
	}
	if data.Username != "" {
		condition.AddCondition("username", "%"+data.Username+"%")
	}
	if data.Name != "" {
		condition.AddCondition("name like ?", "%"+data.Name+"%")
	}
	if data.TagsID > 0 {
		condition.AddCondition("tags_id", data.TagsID)
	}
	if data.IsRecommend > 0 {
		condition.AddCondition("is_recommend = ?", data.IsRecommend)
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
	var imageResponse []dro.ImagesResponse
	if err = utils.ResponseJsonUnmarshal(images, &imageResponse); err != nil {
		return response.ApiError(response.ACCESSERROR, err)
	}

	return response.ApiPageSuccess(imageResponse, total, page, pageSize, total/int64(pageSize) > int64(page))
}

func (is *ImageServer) UpdateServer(di interface{}) *response.ApiResponse {
	data, ok := di.(*dto.ImageUpdate)
	if !ok {
		return response.ApiError(response.ACCESSERROR, nil)
	}

	condition := map[string]interface{}{"id": data.CID}
	total, err := dao.NewCategoryDao().FindByTotal(condition)
	if err != nil {
		return response.ApiError(response.ACCESSERROR, err)
	}
	if data.CID > 0 && total <= 0 {
		return response.ApiError(response.NotFoundCategory, err)
	}

	err = dao.NewImagesDao().UpdateImage(data)
	if err != nil {
		return response.ApiError(response.SaveCategoryErr, err)
	}
	return response.ApiSuccess(nil)
}
func (is *ImageServer) CreateServer(di interface{}) *response.ApiResponse {
	data, ok := di.(*dto.ImageSave)
	if !ok {
		return response.ApiError(response.ACCESSERROR, nil)
	}

	condition := map[string]interface{}{"id": data.CID}
	total, err := dao.NewCategoryDao().FindByTotal(condition)
	if err != nil {
		return response.ApiError(response.ACCESSERROR, err)
	}
	if data.CID > 0 && total <= 0 {
		return response.ApiError(response.NotFoundCategory, err)
	}

	err = dao.NewImagesDao().SaveImage(data)
	if err != nil {
		return response.ApiError(response.SaveCategoryErr, err)
	}
	return response.ApiSuccess(nil)
}
func (is *ImageServer) DeleteServer(di interface{}) *response.ApiResponse {
	return response.ApiSuccess(nil)
}
