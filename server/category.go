package server

import (
	"github.com/gin-gonic/gin"
	"github.com/miaoming3/wallpaper/controller/dto"
	"github.com/miaoming3/wallpaper/dao"
	"github.com/miaoming3/wallpaper/models"
	"github.com/miaoming3/wallpaper/response"
	dto2 "github.com/miaoming3/wallpaper/response/dto"
	"strconv"
)

type CategoryServer struct {
}

func NewCategoryServer() *CategoryServer {
	return &CategoryServer{}
}

func (cg *CategoryServer) IndexServer(c *gin.Context, data *dto.CategoryIndex) *response.ApiResponse {
	condition := map[string]interface{}{
		"status": "1",
	}
	if data.Name != "" {
		condition["name LIKE ?"] = "%" + data.Name + "%"
	}
	if data.Status != "" {
		condition["status"] = data.Status
	}
	pid := 0
	if data.Pid != "" {
		condition["pid"] = data.Pid
		pid, _ = strconv.Atoi(data.Pid)
	}
	categories, err := dao.NewCategoryDao().FindByAll(condition, 0, 0)
	if err != nil {
		return response.ApiError(response.ACCESSERROR, err)
	}
	total, err := dao.NewCategoryDao().FindByTotal(condition)
	if err != nil {
		return response.ApiError(response.ACCESSERROR, err)
	}
	return response.ApiPageSuccess(treeCategory(categories, uint(pid)), total, 0, 0, false)

}

func (cg *CategoryServer) UpdateServer(c *gin.Context, data *dto.UpdateCategory) *response.ApiResponse {
	if data.Pid == data.ID {
		return response.ApiError(response.CategoryParentError, nil)
	}
	condition := map[string]interface{}{"id": data.ID}
	total, err := dao.NewCategoryDao().FindByTotal(condition)
	if err != nil {
		return response.ApiError(response.ACCESSERROR, err)
	}
	if total <= 0 {
		return response.ApiError(response.NotFoundError, err)
	}

	ok, err := dao.NewCategoryDao().UniqueFiled("name", data.Name, data.ID)
	if err != nil {
		return response.ApiError(response.ACCESSERROR, err)
	}
	if !ok {
		return response.ApiError(response.CategoryNameErr, err)
	}

	err = dao.NewCategoryDao().UpdateCategory(data)
	if err != nil {
		return response.ApiError(response.SaveCategoryErr, err)
	}
	return response.ApiSuccess(nil)

}

func (cg *CategoryServer) CreateServer(c *gin.Context, data *dto.SaveCategory) *response.ApiResponse {
	ok, err := dao.NewCategoryDao().UniqueFiled("name", data.Name, 0)
	if err != nil {
		return response.ApiError(response.ACCESSERROR, err)
	}
	if !ok {
		return response.ApiError(response.CategoryNameErr, err)
	}
	condition := map[string]interface{}{"id": data.Pid}
	total, err := dao.NewCategoryDao().FindByTotal(condition)
	if err != nil {
		return response.ApiError(response.ACCESSERROR, err)
	}
	if data.Pid > 0 && total <= 0 {
		return response.ApiError(response.NotFoundPid, err)
	}

	err = dao.NewCategoryDao().SaveCategory(data)
	if err != nil {
		return response.ApiError(response.SaveCategoryErr, err)
	}
	return response.ApiSuccess(nil)
}

func (cg *CategoryServer) DeleteServer(c *gin.Context, data *dto.DeleteCategory) *response.ApiResponse {
	condition := map[string]interface{}{"pid": data.ID}
	total, err := dao.NewCategoryDao().FindByTotal(condition)
	if err != nil {
		return response.ApiError(response.ACCESSERROR, err)
	}
	if total > 0 {
		return response.ApiError(response.FoundSUCCESS, err)
	}
	condition = map[string]interface{}{"id": data.ID}
	total, err = dao.NewCategoryDao().FindByTotal(condition)
	if err != nil {
		return response.ApiError(response.ACCESSERROR, err)
	}
	if total <= 0 {
		return response.ApiError(response.NotFoundError, err)
	}

	if err = dao.NewCategoryDao().DeleteCategory(data); err != nil {
		return response.ApiError(response.DeleteError, err)
	}

	return response.ApiSuccess(nil)

}

func treeCategory(categories []models.Category, pid uint) []*dto2.CategoryListResponse {
	var treeCategories []*dto2.CategoryListResponse
	for _, category := range categories {
		if category.Pid == pid {
			data := &dto2.CategoryListResponse{
				ID:      category.ID,
				Name:    category.Name,
				Pid:     category.Pid,
				Status:  category.Status,
				Sort:    category.Sort,
				Childer: treeCategory(categories, category.ID),
			}
			treeCategories = append(treeCategories, data)
		}
	}
	return treeCategories
}
