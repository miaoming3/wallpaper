package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/miaoming3/wallpaper/app/core/server"
	"github.com/miaoming3/wallpaper/app/message"
	"github.com/miaoming3/wallpaper/app/message/dto"
	"github.com/miaoming3/wallpaper/app/response"
)

type AdminController struct {
}

func NewAdminController() *AdminController {
	return &AdminController{}
}

func (admin *AdminController) Login(c *gin.Context) {
	var login dto.AdminLogin
	if err := c.ShouldBind(&login); err != nil {
		response.Response(c, response.ApiError(message.CLIENTERROR, err))
		return
	}
	response.Response(c, server.NewAdminServer().Login(c, &login))
}

func (admin *AdminController) Info(c *gin.Context) {
	response.Response(c, server.NewAdminServer().Info(c))
}

func (admin *AdminController) ChangePassword(c *gin.Context) {
	var changeData dto.ChangePassword

	if err := c.ShouldBind(&changeData); err != nil {
		response.Response(c, response.ApiError(message.CLIENTERROR, err))
		return
	}
	response.Response(c, server.NewAdminServer().ChangePassword(c, &changeData))
}

func (admin *AdminController) Index(c *gin.Context) {
	var adminSearch dto.AdminSearch
	if err := c.ShouldBind(&adminSearch); err != nil {
		response.Response(c, response.ApiError(message.CLIENTERROR, err))
		return
	}
	response.Response(c, server.NewAdminServer().List(c, &adminSearch))
}
func (admin *AdminController) Created(c *gin.Context) {

}

func (admin *AdminController) ChangeInfo(c *gin.Context) {
	var changeInfo dto.ChangeAdminInfo

	if err := c.ShouldBind(&changeInfo); err != nil {
		response.Response(c, response.ApiError(message.CLIENTERROR, err))
		return
	}
	response.Response(c, server.NewAdminServer().ChangeInfo(c, &changeInfo))
}

func (admin *AdminController) Updated(c *gin.Context) {

}
func (admin *AdminController) Delete(c *gin.Context) {

}
func (admin *AdminController) Logout(c *gin.Context) {
	response.Response(c, server.NewAdminServer().Logout(c))
}
