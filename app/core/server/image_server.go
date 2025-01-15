package server

import (
	"github.com/gin-gonic/gin"
	"github.com/miaoming3/wallpaper/app/core/dao"
	"github.com/miaoming3/wallpaper/app/dto"
	"github.com/miaoming3/wallpaper/app/response"
	response2 "github.com/miaoming3/wallpaper/app/response"
	"github.com/miaoming3/wallpaper/app/response/dro"
	"github.com/miaoming3/wallpaper/app/utils"
)

type ImageServer struct {
}

func NewImageServer() BaseServiceInterface {
	return &ImageServer{}
}

func (is *ImageServer) IndexServer(c *gin.Context, di interface{}) *response.ApiResponse {
	data, ok := di.(*dto.ImageSearch)
	if !ok {
		return response2.ApiError(response2.ACCESSERROR, nil)
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
		return response2.ApiError(response2.ACCESSERROR, err)
	}
	total, err := dao.NewImagesDao().FindByTotal(condition)
	if err != nil {
		return response2.ApiError(response2.ACCESSERROR, err)
	}
	var imageResponse []dro.ImagesResponse
	if err = utils.ResponseJsonUnmarshal(images, &imageResponse); err != nil {
		return response2.ApiError(response2.ACCESSERROR, err)
	}

	return response2.ApiPageSuccess(imageResponse, total, page, pageSize, total/int64(pageSize) > int64(page))
}

func (is *ImageServer) UpdateServer(c *gin.Context, di interface{}) *response.ApiResponse {
	data, ok := di.(*dto.ImageUpdate)
	if !ok {
		return response2.ApiError(response2.ACCESSERROR, nil)
	}
	condition := map[string]interface{}{"id": data.ID}
	total, err := dao.NewImagesDao().FindByTotal(condition)

	if err != nil {
		return response2.ApiError(response2.ACCESSERROR, err)
	}

	if total <= 0 {
		return response2.ApiError(response2.NotFoundImages, err)
	}
	condition = map[string]interface{}{"id": data.CID}
	total, err = dao.NewCategoryDao().FindByTotal(condition)
	if err != nil {
		return response2.ApiError(response2.ACCESSERROR, err)
	}

	if total <= 0 {
		return response2.ApiError(response2.NotFoundCategory, err)
	}

	err = dao.NewImagesDao().UpdateImage(data)
	if err != nil {
		return response2.ApiError(response2.SaveCategoryErr, err)
	}
	return response2.ApiSuccess(nil)
}
func (is *ImageServer) CreateServer(c *gin.Context, di interface{}) *response.ApiResponse {
	data, ok := di.(*dto.ImageSave)
	if !ok {
		return response2.ApiError(response2.ACCESSERROR, nil)
	}

	condition := map[string]interface{}{"id": data.CID}
	total, err := dao.NewCategoryDao().FindByTotal(condition)
	if err != nil {
		return response2.ApiError(response2.ACCESSERROR, err)
	}
	if data.CID > 0 && total <= 0 {
		return response2.ApiError(response2.NotFoundCategory, err)
	}

	err = dao.NewImagesDao().SaveImage(data, uint(c.GetInt("user_uid")))
	if err != nil {
		return response2.ApiError(response2.SaveCategoryErr, err)
	}
	return response2.ApiSuccess(nil)
}
func (is *ImageServer) DeleteServer(di interface{}) *response.ApiResponse {
	return response2.ApiSuccess(nil)
}
