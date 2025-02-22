package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/miaoming3/wallpaper/app/dao"
	"github.com/miaoming3/wallpaper/app/dro"
	"github.com/miaoming3/wallpaper/app/dto"
	"github.com/miaoming3/wallpaper/app/global"
	"github.com/miaoming3/wallpaper/app/message"
	"github.com/miaoming3/wallpaper/app/models"
	"github.com/miaoming3/wallpaper/app/response"
	"github.com/miaoming3/wallpaper/app/utils"
	"strconv"
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
	condition := map[string]interface{}{
		"id":     c.GetInt("uid"),
		"status": 1,
	}
	adminModel, err := dao.NewAdminDao().FindById(condition)
	if err != nil {
		return response.ApiError(message.NotFoundError, err)
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
	condition := map[string]interface{}{
		"id":     uid,
		"status": 1,
	}
	adminModel, err := dao.NewAdminDao().FindById(condition)
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
	uid := c.GetInt("uid")
	condition := map[string]interface{}{
		"id":     uid,
		"status": 1,
	}
	adminModel, err := dao.NewAdminDao().FindById(condition)
	if err != nil {
		return response.ApiError(message.NotFoundError, err)
	}
	if ok, err := admin.UniqueFindFiled(data, uid); !ok {
		return err
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

func (admin *AdminServer) List(c *gin.Context, data *dto.AdminSearch) *response.APi {
	pageSize := c.GetInt("pageSize")
	page := c.GetInt("page")
	condition := map[string]interface{}{
		"id != ?": 1,
	}
	if data.Status != "" {
		condition["status = ?"] = data.Status
	}
	if data.Username != "" {
		condition["username like ?"] = fmt.Sprintf("%v%%", data.Username)
	}
	if data.Email != "" {
		condition["email like ?"] = fmt.Sprintf("%v%%", data.Email)
	}
	if data.Phone != "" {
		condition["phone like ?"] = fmt.Sprintf("%v%%", data.Phone)
	}
	lists, err := dao.NewAdminDao().GetList(condition, page, pageSize)
	if err != nil {
		return response.ApiError(message.NotFoundError, err)
	}
	total, _ := dao.NewAdminDao().GetTotal(condition)
	var adminLists []*dro.AdminInfo
	if len(lists) > 0 {
		for k, v := range lists {
			_ = k
			adminLists = append(adminLists, &dro.AdminInfo{
				Id:        v.ID,
				Status:    v.StatusString(v.Status),
				Password:  v.Password,
				Username:  v.Username,
				Email:     v.Email,
				Phone:     v.Phone,
				ImgUrl:    utils.RemoteUrl(c, v.Avatar),
				Avatar:    v.Avatar,
				CreatedAt: v.CreatedAt.Format(time.DateTime),
			})
		}
	}
	return response.ApiPageSuccess(adminLists, total, page, pageSize, total/int64(pageSize) < int64(page))
}

func (admin *AdminServer) Logout(c *gin.Context) *response.APi {
	_, _ = global.RedisClien.Del(c, fmt.Sprintf("admin:token:%v", c.GetInt("uid"))).Result()
	return response.ApiSuccess(nil)
}

func (admin *AdminServer) Del(c *gin.Context, data *dto.AdminDel) *response.APi {
	condition := map[string]interface{}{
		"id": c.GetInt("uid"),
	}
	adminModel, err := dao.NewAdminDao().FindById(condition)
	if err != nil || c.GetInt("uid") == data.ID || data.ID == 1 {
		return response.ApiError(message.NotFoundError, err)
	}
	if err = dao.NewAdminDao().DeleteRow(data.ID, true, adminModel); err != nil {
		return response.ApiError(message.DeleteError, err)
	}
	return response.ApiSuccess(nil)
}

func (admin *AdminServer) AdminInfo(c *gin.Context, data *dto.AdminDel) *response.APi {
	if c.GetInt("uid") != 1 {
		return response.ApiError(message.NotFoundError, nil)
	}

	condition := map[string]interface{}{
		"id": data.ID,
	}
	adminModel, err := dao.NewAdminDao().FindById(condition)
	if err != nil {
		return response.ApiError(message.NotFoundError, err)
	}
	return response.ApiSuccess(dro.AdminInfo{
		Id:        adminModel.ID,
		Status:    strconv.Itoa(int(adminModel.Status)),
		Password:  adminModel.Password,
		Username:  adminModel.Username,
		Email:     adminModel.Email,
		Phone:     adminModel.Phone,
		ImgUrl:    utils.RemoteUrl(c, adminModel.Avatar),
		Avatar:    adminModel.Avatar,
		CreatedAt: adminModel.CreatedAt.Format(time.DateTime),
	})
}

func (admin *AdminServer) UniqueFindFiled(data interface{}, id int) (bool, *response.APi) {

	return true, response.ApiError(message.UniqueFindFiledError, nil)

}

func (admin *AdminServer) Update(c *gin.Context, data *dto.AdminUpdate) *response.APi {
	uid := c.GetInt("uid")
	if (uid != 1 && uid != data.ID) || (data.ID == 1) {
		return response.ApiError(message.NotFoundError, nil)
	}
	condition := map[string]interface{}{
		"id": data.ID,
	}
	adminModel, err := dao.NewAdminDao().FindById(condition)
	if err != nil {
		return response.ApiError(message.NotFoundError, err)
	}
	if ok, err := admin.UniqueFindFiled(data, data.ID); !ok {
		return err
	}
	if !utils.ComparePoserPassword(adminModel.Password, data.Password) {
		adminModel.Password, _ = utils.GeneratePassword(data.Password)
	}
	adminModel.Username = data.Username
	adminModel.Email = data.Email
	adminModel.Phone = data.Phone
	adminModel.Status = adminModel.StatusUint(data.Status)

	if err = dao.NewAdminDao().UpdateCols(adminModel); err != nil {
		return response.ApiError(message.AdminChangeError, err)
	}

	return response.ApiSuccess(nil)
}

func (admin *AdminServer) Created(c *gin.Context, data *dto.AdminCreated) *response.APi {
	if ok, err := admin.UniqueFindFiled(data, 0); !ok {
		return err
	}
	adminModel := &models.AdminModel{
		Username: data.Username,
		Phone:    data.Phone,
		Email:    data.Email,
		Avatar:   data.Avatar,
	}
	adminModel.Password, _ = utils.GeneratePassword(data.Password)
	adminModel.Status = adminModel.StatusUint(data.Status)
	if err := dao.NewAdminDao().Create(&adminModel).Error; err != nil {
		return response.ApiError(message.CreatedError, err)
	}

	return response.ApiSuccess(nil)
}
