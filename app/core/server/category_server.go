package server

import (
	"github.com/miaoming3/wallpaper/app/core/dao"
	"github.com/miaoming3/wallpaper/app/core/models"
	response2 "github.com/miaoming3/wallpaper/app/response"
	"github.com/miaoming3/wallpaper/app/response/dro"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/miaoming3/wallpaper/app/dto"
	"github.com/miaoming3/wallpaper/app/response"
)

type CategoryServer struct {
}

func NewCategoryServer() BaseServiceInterface {
	return &CategoryServer{}
}

func (cs *CategoryServer) IndexServer(c *gin.Context, di interface{}) *response.APi {
	data, ok := di.(*dto.CategoryIndex)
	if !ok {
		return response2.ApiError(response2.ACCESSERROR, nil)
	}

	condition := dao.NewQueryOption()

	if data.Name != "" {
		condition.AddCondition("name LIKE ? ", "%"+data.Name+"%")
	}
	if data.Status == "" {
		data.Status = "1"
	}
	condition.AddCondition("status = ?", data.Status)
	pid := 0
	if data.Pid != "" {
		condition.AddCondition("pid = ?", data.Pid)
		pid, _ = strconv.Atoi(data.Pid)
	}

	categories, err := dao.NewCategoryDao().FindByAll(condition)
	if err != nil {
		return response2.ApiError(response2.ACCESSERROR, err)
	}
	total, err := dao.NewCategoryDao().FindByTotal(condition)
	if err != nil {
		return response2.ApiError(response2.ACCESSERROR, err)
	}
	return response2.ApiPageSuccess(treeCategory(categories, uint(pid)), total, 0, 0, false)

}

func (cs *CategoryServer) UpdateServer(c *gin.Context, di interface{}) *response.APi {

	data, ok := di.(*dto.UpdateCategory)
	if !ok {
		return response2.ApiError(response2.ACCESSERROR, nil)
	}
	if data.Pid == data.ID {
		return response2.ApiError(response2.CategoryParentError, nil)
	}
	condition := map[string]interface{}{"id": data.ID}
	total, err := dao.NewCategoryDao().FindByTotal(condition)
	if err != nil {
		return response2.ApiError(response2.ACCESSERROR, err)
	}
	if total <= 0 {
		return response2.ApiError(response2.NotFoundError, err)
	}

	ok, err = dao.NewCategoryDao().UniqueFiled("name", data.Name, data.ID)
	if err != nil {
		return response2.ApiError(response2.ACCESSERROR, err)
	}
	if !ok {
		return response2.ApiError(response2.CategoryNameErr, err)
	}

	err = dao.NewCategoryDao().UpdateCategory(data)
	if err != nil {
		return response2.ApiError(response2.SaveCategoryErr, err)
	}
	return response2.ApiSuccess(nil)

}

func (cs *CategoryServer) CreateServer(c *gin.Context, di interface{}) *response.APi {
	data, ok := di.(*dto.SaveCategory)
	if !ok {
		return response2.ApiError(response2.ACCESSERROR, nil)
	}
	ok, err := dao.NewCategoryDao().UniqueFiled("name", data.Name, 0)
	if err != nil {
		return response2.ApiError(response2.ACCESSERROR, err)
	}
	if !ok {
		return response2.ApiError(response2.CategoryNameErr, err)
	}
	condition := map[string]interface{}{"id": data.Pid}
	total, err := dao.NewCategoryDao().FindByTotal(condition)
	if err != nil {
		return response2.ApiError(response2.ACCESSERROR, err)
	}
	if data.Pid > 0 && total <= 0 {
		return response2.ApiError(response2.NotFoundPid, err)
	}

	err = dao.NewCategoryDao().SaveCategory(data)
	if err != nil {
		return response2.ApiError(response2.SaveCategoryErr, err)
	}
	return response2.ApiSuccess(nil)
}

func (cs *CategoryServer) DeleteServer(di interface{}) *response.APi {
	data, ok := di.(*dto.DeleteCategory)
	if !ok {
		return response2.ApiError(response2.ACCESSERROR, nil)
	}
	condition := map[string]interface{}{"pid": data.ID}
	total, err := dao.NewCategoryDao().FindByTotal(condition)
	if err != nil {
		return response2.ApiError(response2.ACCESSERROR, err)
	}
	if total > 0 {
		return response2.ApiError(response2.FoundSUCCESS, err)
	}
	condition = map[string]interface{}{"id": data.ID}
	total, err = dao.NewCategoryDao().FindByTotal(condition)
	if err != nil {
		return response2.ApiError(response2.ACCESSERROR, err)
	}
	if total <= 0 {
		return response2.ApiError(response2.NotFoundError, err)
	}

	if err = dao.NewCategoryDao().DeleteCategory(data); err != nil {
		return response2.ApiError(response2.DeleteError, err)
	}

	return response2.ApiSuccess(nil)

}

func treeCategory(categories []models.Category, pid uint) []*dro.CategoryListResponse {
	var treeCategories []*dro.CategoryListResponse
	for _, category := range categories {
		if category.Pid == pid {
			data := &dro.CategoryListResponse{
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
