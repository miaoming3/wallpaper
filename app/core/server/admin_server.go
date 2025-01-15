package server

import (
	"github.com/miaoming3/wallpaper/app/core/dao"
	"github.com/miaoming3/wallpaper/app/core/models"
	"github.com/miaoming3/wallpaper/app/global"
	response2 "github.com/miaoming3/wallpaper/app/response"
	"github.com/miaoming3/wallpaper/app/response/dro"
	"github.com/miaoming3/wallpaper/app/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/miaoming3/wallpaper/app/dto"
	"reflect"

	"github.com/miaoming3/wallpaper/app/response"
	"golang.org/x/crypto/bcrypt"
)

type AdminServer struct {
}

func NewAdminServer() *AdminServer {
	return &AdminServer{}
}

func checkCaptcha(id string, captcha string) (bool, error) {
	if global.SysConfig.Model == "debug" && captcha != "8888" && id != "123456789" {
		return false, nil
	}

	return true, nil
}

func checkPassword(passwordHash string, password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password)); err != nil {
		return false
	}
	return true
}

func (as *AdminServer) LoginServer(c *gin.Context, data *dto.AdminLoginData) *response.ApiResponse {
	ok, err := checkCaptcha(data.CaptchaID, data.Captcha)
	if err != nil || !ok {
		return response2.ApiError(response2.CAPTCHAERROR, err)
	}

	admin, err := dao.NewAdminDao().FindByUsername(data.Username)
	if err != nil || reflect.DeepEqual(admin, models.Admin{}) || !checkPassword(admin.Password, data.Password) {
		return response2.ApiError(response2.ADMINORPASSWORD, err)
	}
	return response2.ApiSuccess(dro.AdminLoginResponse{
		Token:      utils.GenerateRandomStringMath(16),
		UID:        admin.Uid,
		Expression: time.Now().Format(time.DateTime),
		Username:   admin.Username,
	})
}

func (as *AdminServer) IndexServer(ctx *gin.Context) ([]interface{}, error) {

	return nil, nil
}
