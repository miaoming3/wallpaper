package server

import (
	"fmt"
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

	fmt.Println(condition)

	categories, err := dao.NewCategoryDao().FindByAll(condition, data.Page, data.PageSize)
	if err != nil {
		return response.ApiError(response.ACCESSERROR, err)
	}
	total, err := dao.NewCategoryDao().FindByTotal(condition)
	if err != nil {
		return response.ApiError(response.ACCESSERROR, err)
	}
	return response.ApiPageSuccess(treeCategory(categories, uint(pid)), total, data.Page, data.PageSize, total/int64(data.PageSize) >= int64(data.Page))

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
