package dao

import (
	"github.com/miaoming3/wallpaper/global"
	"github.com/miaoming3/wallpaper/models"
	"gorm.io/gorm"
)

type ImagesDao struct {
	*gorm.DB
}

func NewImagesDao() *ImagesDao {
	return &ImagesDao{global.DbClient}
}

func (dao *ImagesDao) FindByAll(condition interface{}, page int, pageSize int) ([]models.Image, error) {
	var images []models.Image

	offset := (page - 1) * pageSize
	if err := dao.Preload("User").Preload("Category").Preload("Tags").Model(models.Image{}).Where(condition).Offset(offset).Limit(pageSize).Find(&images).Error; err != nil {
		return nil, err
	}
	return images, nil
}

func (dao *ImagesDao) FindByTotal(conditions interface{}) (total int64, err error) {
	if err = dao.Preload("users").Preload("Category").Preload("image_tags").Model(models.Image{}).Where(conditions).Count(&total).Error; err != nil {
		return 0, err
	}
	return
}
