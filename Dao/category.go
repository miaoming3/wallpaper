package dao

import (
	"github.com/miaoming3/wallpaper/global"
	"github.com/miaoming3/wallpaper/models"
	"gorm.io/gorm"
)

type CategoryDao struct {
	*gorm.DB
}

func NewCategoryDao() *CategoryDao {
	return &CategoryDao{global.DbClient}
}

func (dao *CategoryDao) FindByAll(condition interface{}, page int, pageSize int) ([]models.Category, error) {
	var categories []models.Category
	if pageSize == 0 {
		pageSize = 20
	}
	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize
	if err := dao.Model(models.Category{}).Where(condition).Offset(offset).Limit(pageSize).Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (dao *CategoryDao) FindByTotal(conditions interface{}) (total int64, err error) {
	if err = dao.Model(models.Category{}).Where(conditions).Count(&total).Error; err != nil {
		return 0, err
	}
	return
}
