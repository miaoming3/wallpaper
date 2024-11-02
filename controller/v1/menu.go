package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/miaoming3/wallpaper/controller/dto"
	"github.com/miaoming3/wallpaper/response"
	"github.com/miaoming3/wallpaper/server"
)

type MenuController struct {
	server.BaseServiceInterface
}

func NewMenuController() *MenuController {
	return &MenuController{server.NewMenuServe()}
}

// Index
// @Summary 菜单列表
// @Tags 菜单管理
// @Description 获取菜单列表
// @Produce application/json
// @Router /menu/index [get]
func (mc *MenuController) Index(c *gin.Context) {
	var index dto.MenuIndex
	if err := c.ShouldBind(&index); err != nil {
		response.Response(c, response.ApiError(response.CLIENTERROR, err))
		return
	}

	response.Response(c, mc.IndexServer(&index))
}

// Save
// @Summary 创建菜单
// @Tags 菜单管理
// @Description 创建菜单
// @Param SaveMenu body dto.SaveMenu true "创建菜单"
// @Router /menu/save [post]
func (mc *MenuController) Save(c *gin.Context) {

	var saveMenu dto.SaveMenu
	if err := c.ShouldBind(&saveMenu); err != nil {
		response.Response(c, response.ApiError(response.CLIENTERROR, err))
		return
	}
	response.Response(c, mc.CreateServer(&saveMenu))
}

// Update
// @Summary 修改菜单
// @Tags 菜单管理
// @Description 修改菜单
// @Param UpdateMenu body dto.UpdateMenu true "修改菜单"
// @Success 200 {object} response.ApiResponse
// @Router /menu/update [put]
func (mc *MenuController) Update(c *gin.Context) {
	var updateMenu dto.UpdateMenu
	if err := c.ShouldBind(&updateMenu); err != nil {
		response.Response(c, response.ApiError(response.CLIENTERROR, err))
		return
	}
	response.Response(c, mc.UpdateServer(&updateMenu))
}

// Delete
// @Summary 删除菜单
// @Tags 菜单管理
// @Description 删除菜单
// @Param  DeleteMenu body dto.DeleteMenu true "删除菜单参数"
// @Router /menu/del [delete]
func (mc *MenuController) Delete(c *gin.Context) {
	var deleteMenu dto.DeleteMenu
	if err := c.ShouldBind(&deleteMenu); err != nil {
		response.Response(c, response.ApiError(response.CLIENTERROR, err))
		return
	}
	response.Response(c, mc.DeleteServer(&deleteMenu))
}
