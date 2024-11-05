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

func (dao *ImagesDao) getModel(condition *QueryOption) *gorm.DB {
	query := dao.Model(models.Image{})
	for key, value := range condition.Conditions {
		if key == "tags_id" && value != nil {
			subQuery := dao.Model(models.ImageTags{}).Select("image_id").Where("tags_id = ?", value)
			query = query.Where("id IN (?)", subQuery)
		} else if key == "username" && value != nil {
			subQuery := dao.Model(models.Users{}).Select("uid").Where("username like ?", value)
			query = query.Where("user_id IN (?)", subQuery)
		} else {
			query = query.Where(key, value)
		}
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

func (dao *ImagesDao) FindByAll(condition *QueryOption, page int, pageSize int) ([]models.Image, error) {
	var images []models.Image
	offset := (page - 1) * pageSize
	if err := dao.getModel(condition).Offset(offset).Limit(pageSize).Find(&images).Error; err != nil {
		return nil, err
	}
	return images, nil
}

func (dao *ImagesDao) FindByTotal(condition interface{}) (total int64, err error) {
	var query *gorm.DB

	if conditions, ok := condition.(*QueryOption); ok {
		query = dao.getModel(conditions)
	} else {
		query = dao.Model(models.Image{}).Where(condition)
	}

	if err = query.Count(&total).Error; err != nil {
		return 0, err
	}
	return
}
