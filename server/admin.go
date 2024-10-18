package server

import (
	"github.com/gin-gonic/gin"
	"github.com/miaoming3/wallpaper/controller/dto"
	"github.com/miaoming3/wallpaper/global"
	dto2 "github.com/miaoming3/wallpaper/response/dto"
	"time"

	"github.com/miaoming3/wallpaper/dao"
	"github.com/miaoming3/wallpaper/models"
	"github.com/miaoming3/wallpaper/response"
	"golang.org/x/crypto/bcrypt"
	"reflect"
)

type AdminServer struct {
}

func NewAdminServer() *AdminServer {
	return &AdminServer{}
}

func checkCaptcha(id string, captcha string) (bool, error) {
	if global.SysConfig.Dev && captcha != "8888" && id != "123456789" {
		return false, nil
	}

	return true, nil
}

func checkPassword(hash_password string, password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hash_password), []byte(password)); err != nil {
		return false
	}
	return true
}
func (as *AdminServer) LoginServer(c *gin.Context, data *dto.AdminLoginData) *response.ApiResponse {
	ok, err := checkCaptcha(data.CaptchaID, data.Captcha)
	if err != nil || !ok {
		return response.ApiError(response.CAPTCHAERROR, err)
	}

	admin, err := dao.NewAdminDao().FindByUsername(data.Username)
	if err != nil || reflect.DeepEqual(admin, models.Admin{}) || !checkPassword(admin.Password, data.Password) {
		return response.ApiError(response.ADMINORPASSWORD, err)
	}

	return response.ApiSuccess(dto2.AdminLoginResponse{
		Token:      "123456",
		UID:        admin.Uid,
		Expression: time.Now().Format(time.DateTime),
		Username:   admin.Username,
	})
}

func (as *AdminServer) IndexServer(ctx *gin.Context) ([]interface{}, error) {

	return nil, nil
}