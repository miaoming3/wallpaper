package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/miaoming3/wallpaper/controller/dto"
	"github.com/miaoming3/wallpaper/response"
	"github.com/miaoming3/wallpaper/server"
)

type CategoryController struct {
	server.CategoryService
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
		response.Response(c, response.ApiError(response.CLIENTERROR, err))
		return
	}

	response.Response(c, cg.IndexServer(&searchData))
}

// Save
// @Summary 创建分类
// @Tags 分类管理
// @Description 创建分类
// @Param saveCategory body dto.SaveCategory true "Category data to save"
// @Router /category/save [post]
func (cg *CategoryController) Save(c *gin.Context) {

	var saveCategory dto.SaveCategory
	if err := c.ShouldBind(&saveCategory); err != nil {
		response.Response(c, response.ApiError(response.CLIENTERROR, err))
		return
	}
	response.Response(c, cg.CreateServer(&saveCategory))
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
		response.Response(c, response.ApiError(response.CLIENTERROR, err))
		return
	}
	response.Response(c, cg.UpdateServer(&updateCategory))
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
		response.Response(c, response.ApiError(response.CLIENTERROR, err))
		return
	}
	response.Response(c, cg.DeleteServer(&deleteCategory))
}
