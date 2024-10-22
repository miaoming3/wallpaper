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

	query := dao.Model(models.Category{}).Where(condition)
	if page > 0 && pageSize > 0 {
		offset := (page - 1) * pageSize
		query.Offset(offset).Limit(pageSize)
	}
	if err := query.Find(&categories).Error; err != nil {
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
