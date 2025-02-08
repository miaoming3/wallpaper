package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/miaoming3/wallpaper/app/core/dao"
	"github.com/miaoming3/wallpaper/app/global"
	"github.com/miaoming3/wallpaper/app/message"
	"github.com/miaoming3/wallpaper/app/message/dro"
	"github.com/miaoming3/wallpaper/app/message/dto"
	"github.com/miaoming3/wallpaper/app/response"
	"github.com/miaoming3/wallpaper/app/utils"
	"time"
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

	if !utils.ComparePoserPassword(adminModel.Password, data.Password) {
		return response.ApiError(message.ADMINORPASSWORD, err)
	}
	token := utils.GenerateRandomStringMath(16)
	err = global.RedisClien.Set(c, fmt.Sprintf("admin:token:%v", adminModel.ID), token, time.Hour*24).Err()
	if err != nil {
		return response.ApiError(message.ACCESSERROR, err)
	}

	return response.ApiSuccess(&dro.LoginResponse{
		Uid:   adminModel.ID,
		Token: token,
	})
}

func (admin *AdminServer) Info(c *gin.Context) *response.APi {

	adminModel, err := dao.NewAdminDao().FindById(c.GetInt("uid"), 1)
	if err != nil {
		return response.ApiError(message.ADMINORPASSWORD, err)
	}
	return response.ApiSuccess(dro.AdminInfo{
		Id:        adminModel.ID,
		Status:    adminModel.StatusString(adminModel.Status),
		Password:  adminModel.Password,
		Username:  adminModel.Username,
		Email:     adminModel.Email,
		Phone:     adminModel.Phone,
		ImgUrl:    utils.RemoteUrl(c, adminModel.Avatar),
		Avatar:    adminModel.Avatar,
		CreatedAt: adminModel.CreatedAt.Format(time.DateTime),
	})
}

func (admin *AdminServer) ChangePassword(c *gin.Context, data *dto.ChangePassword) *response.APi {
	uid := c.GetInt("uid")
	adminModel, err := dao.NewAdminDao().FindById(uid, 1)
	if err != nil {
		return response.ApiError(message.NotFoundRow, err)
	}
	if !utils.ComparePoserPassword(adminModel.Password, data.OldPassword) {
		return response.ApiError(message.PasswordError, err)
	}

	password, _ := utils.GeneratePassword(data.Password)
	if err = dao.NewAdminDao().UpdateCol("password", password, adminModel); err != nil {
		return response.ApiError(message.AdminChangeError, err)
	}
	global.RedisClien.Del(c, fmt.Sprintf("admin:token:%v", uid))
	return response.ApiSuccess(nil)
}

func (admin *AdminServer) ChangeInfo(c *gin.Context, data *dto.ChangeAdminInfo) *response.APi {
	adminModel, err := dao.NewAdminDao().FindById(c.GetInt("uid"), 1)
	if err != nil {
		return response.ApiError(message.ADMINORPASSWORD, err)
	}
	adminModel.Username = data.Username
	adminModel.Email = data.Email
	adminModel.Phone = data.Phone
	if data.Avatar != "" {
		adminModel.Avatar = data.Avatar
	}
	if err = dao.NewAdminDao().UpdateCols(adminModel); err != nil {
		return response.ApiError(message.AdminChangeError, err)
	}
	return response.ApiSuccess(nil)
}

func (admin *AdminServer) Logout(c *gin.Context) *response.APi {
	_, _ = global.RedisClien.Del(c, fmt.Sprintf("admin:token:%v", c.GetInt("uid"))).Result()
	return response.ApiSuccess(nil)
}
