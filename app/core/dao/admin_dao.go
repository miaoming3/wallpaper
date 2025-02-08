package dao

import (
	"github.com/miaoming3/wallpaper/app/global"
	"github.com/miaoming3/wallpaper/app/models"
	"gorm.io/gorm"
)

type AdminDao struct {
	*gorm.DB
}

func NewAdminDao() *AdminDao {
	return &AdminDao{global.DbClient}
}

func (dao *AdminDao) FindByName(username string, status int) (admin *models.AdminModel, err error) {
	if err = dao.Model(&models.AdminModel{}).Where("username =? and status= ?", username, status).First(&admin).Error; err != nil {
		return &models.AdminModel{}, err
	}
	return
}

func (dao *AdminDao) FindById(uid int, status int) (admin *models.AdminModel, err error) {

	if err = dao.Model(&models.AdminModel{}).Where("id = ? and status = ?", uid, status).First(&admin).Error; err != nil {
		return &models.AdminModel{}, err
	}
	return
}

func (dao *AdminDao) UpdateCol(col string, val interface{}, adminModel *models.AdminModel) error {
	return dao.Model(&models.AdminModel{}).Where("id =? and status= ?", adminModel.ID, adminModel.Status).Update(col, val).Error
}

func (dao *AdminDao) UpdateCols(admin *models.AdminModel) error {
	return dao.Model(&models.AdminModel{}).Where("id = ?", admin.ID).Updates(&admin).Error
}
