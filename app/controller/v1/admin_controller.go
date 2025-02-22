package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/miaoming3/wallpaper/app/dto"
	"github.com/miaoming3/wallpaper/app/message"
	"github.com/miaoming3/wallpaper/app/response"
	"github.com/miaoming3/wallpaper/app/server"
)

type AdminController struct {
}

func NewAdminController() *AdminController {
	return &AdminController{}
}

// Login  管理员登录
func (admin *AdminController) Login(c *gin.Context) {
	var login dto.AdminLogin
	if err := c.ShouldBind(&login); err != nil {
		response.Response(c, response.ApiError(message.CLIENTERROR, err))
		return
	}
	response.Response(c, server.NewAdminServer().Login(c, &login))
}

// Info  管理员信息
func (admin *AdminController) Info(c *gin.Context) {
	response.Response(c, server.NewAdminServer().Info(c))
}

// ChangePassword 修改密码
func (admin *AdminController) ChangePassword(c *gin.Context) {
	var changeData dto.ChangePassword

	if err := c.ShouldBind(&changeData); err != nil {
		response.Response(c, response.ApiError(message.CLIENTERROR, err))
		return
	}
	response.Response(c, server.NewAdminServer().ChangePassword(c, &changeData))
}

// Index 管理员列表
func (admin *AdminController) Index(c *gin.Context) {
	var adminSearch dto.AdminSearch
	if err := c.ShouldBind(&adminSearch); err != nil {
		response.Response(c, response.ApiError(message.CLIENTERROR, err))
		return
	}
	response.Response(c, server.NewAdminServer().List(c, &adminSearch))
}
func (admin *AdminController) Created(c *gin.Context) {
	var adminData dto.AdminCreated

	if err := c.ShouldBind(&adminData); err != nil {
		response.Response(c, response.ApiError(message.CLIENTERROR, err))
		return
	}
	response.Response(c, server.NewAdminServer().Created(c, &adminData))
}

// ChangeInfo 修改个人资料
func (admin *AdminController) ChangeInfo(c *gin.Context) {
	var changeInfo dto.ChangeAdminInfo

	if err := c.ShouldBind(&changeInfo); err != nil {
		response.Response(c, response.ApiError(message.CLIENTERROR, err))
		return
	}
	response.Response(c, server.NewAdminServer().ChangeInfo(c, &changeInfo))
}

func (admin *AdminController) Updated(c *gin.Context) {
	var adminData dto.AdminUpdate

	if err := c.ShouldBind(&adminData); err != nil {
		response.Response(c, response.ApiError(message.CLIENTERROR, err))
		return
	}
	response.Response(c, server.NewAdminServer().Update(c, &adminData))
}

// Delete  删除管理员
func (admin *AdminController) Delete(c *gin.Context) {
	var delData dto.AdminDel

	if err := c.ShouldBind(&delData); err != nil {
		response.Response(c, response.ApiError(message.CLIENTERROR, err))
		return
	}
	response.Response(c, server.NewAdminServer().Del(c, &delData))
}

// Logout  退出
func (admin *AdminController) Logout(c *gin.Context) {
	response.Response(c, server.NewAdminServer().Logout(c))
}

func (admin *AdminController) GetAdminInfo(c *gin.Context) {
	var adminInfo dto.AdminDel
	if err := c.ShouldBind(&adminInfo); err != nil {
		response.Response(c, response.ApiError(message.CLIENTERROR, err))
		return
	}
	response.Response(c, server.NewAdminServer().AdminInfo(c, &adminInfo))
}
