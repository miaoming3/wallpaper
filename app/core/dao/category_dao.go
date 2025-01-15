package dao

import (
	"fmt"
	"github.com/miaoming3/wallpaper/app/core/models"
	"github.com/miaoming3/wallpaper/app/global"

	"github.com/miaoming3/wallpaper/app/dto"
	"gorm.io/gorm"
)

type CategoryDao struct {
	*gorm.DB
}

func NewCategoryDao() *CategoryDao {
	return &CategoryDao{global.DbClient}
}

func (dao *CategoryDao) getDBX(condition *QueryOption) *gorm.DB {
	query := dao.Model(models.Category{})
	for key, value := range condition.Conditions {
		query = query.Where(key, value)
	}
	for preloadKey, preloadCondition := range condition.Preloads {
		if preloadCondition != nil {
			query = query.Preload(preloadKey, preloadCondition)
		} else {
			query = query.Preload(preloadKey)
		}
	}
	return query
}

func (dao *CategoryDao) FindByAll(condition *QueryOption) ([]models.Category, error) {
	var categories []models.Category
	if err := dao.getDBX(condition).Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (dao *CategoryDao) FindByTotal(condition interface{}) (total int64, err error) {
	var query *gorm.DB
	if conditions, ok := condition.(*QueryOption); ok {
		query = dao.getDBX(conditions)
	} else {
		query = dao.Model(models.Category{}).Where(condition)
	}
	if err = query.Count(&total).Error; err != nil {
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
