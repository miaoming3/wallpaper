package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/miaoming3/wallpaper/app/core/server"
	"github.com/miaoming3/wallpaper/app/dto"
	response2 "github.com/miaoming3/wallpaper/app/response"
)

type CategoryController struct {
	server.BaseServiceInterface
}

// NewCategoryController
// @Tags 分类管理
// @Description  管理分类
// @Schemes http, https
// @Success 200 {object} response.ApiResponse
func NewCategoryController() *CategoryController {
	return &CategoryController{server.NewCategoryServer()}
}

// Index
// @Summary 分类列表
// @Tags 分类管理
// @Description 获取分类列表
// @Produce application/json
// @Router /category/index [get]
func (cg *CategoryController) Index(c *gin.Context) {
	var searchData dto.CategoryIndex
	if err := c.ShouldBind(&searchData); err != nil {
		response2.Response(c, response2.ApiError(response2.CLIENTERROR, err))
		return
	}

	response2.Response(c, cg.IndexServer(c, &searchData))
}

// Save
// @Summary 创建分类
// @Tags 分类管理
// @Description 创建分类
// @Accept json,application/x-www-form-urlencoded
// @Produce application/json
// @Param saveCategory body dto.SaveCategory true "保存分类参数" // JSON 请求体
// @Router /category/save [post]
func (cg *CategoryController) Save(c *gin.Context) {
	var saveCategory dto.SaveCategory
	if err := c.ShouldBind(&saveCategory); err != nil {
		response2.Response(c, response2.ApiError(response2.CLIENTERROR, err))
		return
	}

	response2.Response(c, cg.CreateServer(c, &saveCategory))
}

// Update
// @Summary 修改分类
// @Tags 分类管理
// @Description 修改分类
// @Param UpdateCategory body dto.UpdateCategory true "Category data to update"
// @Success 200 {object} response.ApiResponse
// @Router /category/update [put]
func (cg *CategoryController) Update(c *gin.Context) {
	var updateCategory dto.UpdateCategory

	if err := c.ShouldBind(&updateCategory); err != nil {
		response2.Response(c, response2.ApiError(response2.CLIENTERROR, err))
		return
	}
	response2.Response(c, cg.UpdateServer(c, &updateCategory))
}

// Delete
// @Summary 删除分类
// @Tags 分类管理
// @Description 删除分类
// @Param DeleteCategory body dto.DeleteCategory true "Category data to delete"
// @Router /category/del [delete]
func (cg *CategoryController) Delete(c *gin.Context) {
	var deleteCategory dto.DeleteCategory
	if err := c.ShouldBind(&deleteCategory); err != nil {
		response2.Response(c, response2.ApiError(response2.CLIENTERROR, err))
		return
	}
	response2.Response(c, cg.DeleteServer(&deleteCategory))
}
