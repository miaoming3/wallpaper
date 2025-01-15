package dao

import (
	"github.com/miaoming3/wallpaper/app/core/models"
	"github.com/miaoming3/wallpaper/app/global"
	"gorm.io/gorm"
)

type AdminDao struct {
	*gorm.DB
}

func NewAdminDao() *AdminDao {
	return &AdminDao{global.DbClient}
}

func (dao *AdminDao) FindByUsername(username string) (admin models.Admin, err error) {
	if err = dao.Model(models.Admin{}).Where("username = ? and status =1", username).First(&admin).Error; err != nil {
		return models.Admin{}, err
	}
	return admin, nil
}
