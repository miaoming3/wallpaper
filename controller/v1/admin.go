package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/miaoming3/wallpaper/controller/dto"
	"github.com/miaoming3/wallpaper/response"
	"github.com/miaoming3/wallpaper/server"
)

type adminController struct {
	Server *server.AdminServer
}

func NewAdminController() *adminController {
	return &adminController{
		Server: server.NewAdminServer(),
	}
}

func (ac adminController) Login(c *gin.Context) {
	var requestLogin dto.AdminLoginData
	if err := c.ShouldBind(&requestLogin); err != nil {
		response.Response(c, response.ApiError(response.CLIENTERROR, err))
		return
	}
	response.Response(c, ac.Server.LoginServer(c, &requestLogin))
}
