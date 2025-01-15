package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/miaoming3/wallpaper/app/core/server"
	"github.com/miaoming3/wallpaper/app/dto"
	response2 "github.com/miaoming3/wallpaper/app/response"
)

type AdminController struct {
	*server.AdminServer
}

func NewAdminController() *AdminController {
	return &AdminController{
		server.NewAdminServer(),
	}
}

// Login
// @Summary 管理员登录
// @Tags 管理员管理
// @Description 管理员用户登录
// @Produce application/json
// @Param AdminLoginData body dto.AdminLoginData true "管理员登录字段"
// @Router /admin/login [post]
func (ac *AdminController) Login(c *gin.Context) {
	var requestLogin dto.AdminLoginData
	if err := c.ShouldBind(&requestLogin); err != nil {
		response2.Response(c, response2.ApiError(response2.CLIENTERROR, err))
		return
	}
	response2.Response(c, ac.LoginServer(c, &requestLogin))
}

func (ac *AdminController) Index(c *gin.Context) {

}
func (ac *AdminController) Save(c *gin.Context) {

}
func (ac *AdminController) Update(c *gin.Context) {

}
func (ac *AdminController) Delete(c *gin.Context) {

}
