package dao

import (
	"fmt"
	"github.com/miaoming3/wallpaper/controller/dto"
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

func (dao *CategoryDao) UniqueFiled(filed string, value interface{}, id uint) (bool, error) {
	query := dao.Model(models.Category{}).Where(fmt.Sprintf("%s = ?", filed), value)
	if id > 0 {
		query.Where("id != ?", id)
	}
	var count int64
	if err := query.Count(&count).Error; err != nil {
		return false, err
	}
	return count == 0, nil
}

func (dao *CategoryDao) SaveCategory(data *dto.SaveCategory) error {
	saveData := &models.Category{
		Name:   data.Name,
		Pid:    data.Pid,
		Sort:   data.Sort,
		Status: data.Status,
	}
	return dao.Model(models.Category{}).Create(saveData).Error
}

func (dao *CategoryDao) UpdateCategory(data *dto.UpdateCategory) error {
	saveData := &models.Category{
		Name:   data.Name,
		Pid:    data.Pid,
		Sort:   data.Sort,
		Status: data.Status,
	}
	return dao.Model(models.Category{}).Where("id = ?", data.ID).Updates(saveData).Error
}

func (dao *CategoryDao) DeleteCategory(data *dto.DeleteCategory) error {
	return dao.Delete(&models.Category{}, data.ID).Error
}
