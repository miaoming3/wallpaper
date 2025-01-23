package server

import (
	"github.com/gin-gonic/gin"
	"github.com/miaoming3/wallpaper/app/core/dao"
	"github.com/miaoming3/wallpaper/app/global"
	"github.com/miaoming3/wallpaper/app/message"
	"github.com/miaoming3/wallpaper/app/message/dro"
	"github.com/miaoming3/wallpaper/app/message/dto"
	"github.com/miaoming3/wallpaper/app/response"
	"github.com/miaoming3/wallpaper/app/utils"
)

type AdminServer struct {
}

func NewAdminServer() *AdminServer {
	return &AdminServer{}
}

func (admin *AdminServer) Login(c *gin.Context, data *dto.AdminLogin) *response.APi {
	if ok := global.Captcha.Verify(data.CaptchaId, data.Captcha, true); !ok {
		return response.ApiError(message.CAPTCHAERROR, nil)
	}

	adminModel, err := dao.NewAdminDao().FindByName(data.Username, 1)
	if err != nil {
		return response.ApiError(message.ADMINORPASSWORD, err)
	}

	if ok := utils.ComparePoserPassword(adminModel.Password, data.Password); ok {
		return response.ApiError(message.ADMINORPASSWORD, err)
	}

	return response.ApiSuccess(&dro.LoginResponse{
		Uid:   adminModel.ID,
		Token: utils.GenerateRandomStringMath(16),
	})
}

func (admin *AdminServer) Info(c *gin.Context) *response.APi {
	dao.NewAdminDao()
	return response.ApiSuccess(nil)
}

func (admin *AdminServer) ChangePassword(c *gin.Context, data *dto.ChangePassword) *response.APi {
	adminModel, err := dao.NewAdminDao().FindById(c.GetUint("uid"), 1)
	if err != nil {
		return response.ApiError(message.ADMINORPASSWORD, err)
	}
	if ok := utils.ComparePoserPassword(adminModel.Password, data.Password); ok {
		return response.ApiError(message.ADMINORPASSWORD, err)
	}

	password, _ := utils.GeneratePassword(data.Password)
	if err = dao.NewAdminDao().UpdateCol("password", password, adminModel); err != nil {
		return response.ApiError(message.ADMINORPASSWORD, err)
	}
	return response.ApiSuccess(nil)
}
