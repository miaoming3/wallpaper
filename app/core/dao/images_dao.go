package dao

import (
	"github.com/miaoming3/wallpaper/app/core/models"
	"github.com/miaoming3/wallpaper/app/dto"
	"github.com/miaoming3/wallpaper/app/global"
	"gorm.io/gorm"
)

type ImagesDao struct {
	*gorm.DB
}

func NewImagesDao() *ImagesDao {
	return &ImagesDao{global.DbClient}
}

// getDBX
func (dao *ImagesDao) getDBX(condition *QueryOption) *gorm.DB {
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

	if err := dao.getDBX(condition).Offset(offset).Limit(pageSize).Find(&images).Error; err != nil {
		return nil, err
	}
	return images, nil
}
func (dao *ImagesDao) SaveImage(data *dto.ImageSave, uid uint) error {
	image := &models.Image{
		Name:        data.Name,
		Type:        data.Type,
		Cid:         data.CID,
		Url:         data.Url,
		Path:        data.Path,
		IsRecommend: data.IsRecommend,
		UserID:      uid,
	}
	return dao.Begin().Transaction(func(db *gorm.DB) error {
		if err := db.Model(models.Image{}).Create(image).Error; err != nil {
			db.Rollback()
			return err
		}
		if len(data.Tags) > 0 {
			var tags []models.Tags
			if err := db.Model(models.Image{}).Where("id IN ?", data.Tags).Find(&tags).Error; err != nil {
				db.Rollback()
				return err
			}

			if err := db.Model(models.Image{}).Model(image).Association("Tags").Append(tags); err != nil {
				db.Rollback()
				return err
			}
		}
		db.Commit()
		return nil
	})

}
func (dao *ImagesDao) UpdateImage(data *dto.ImageUpdate) error {
	// 创建Image实例
	image := &models.Image{
		Name:        data.Name,
		Type:        data.Type,
		Cid:         data.CID,
		Url:         data.Url,
		Path:        data.Path,
		IsRecommend: data.IsRecommend,
	}
	return dao.Begin().Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(models.Image{}).Where("id = ?", data.ID).Updates(image).Error; err != nil {
			tx.Rollback()
			return err
		}
		if len(data.Tags) > 0 {
			var tags []models.Tags
			if err := tx.Model(models.Image{}).Where("id IN ?", data.Tags).Find(&tags).Error; err != nil {
				tx.Rollback()
				return err
			}
			// 使用 Replace 关联查询出来的 Tags 实例
			if err := tx.Model(&models.Image{Model: gorm.Model{ID: data.ID}}).Association("Tags").Replace(tags); err != nil {
				tx.Rollback()
				return err
			}
		}
		tx.Commit()
		return nil
	})

}
func (dao *ImagesDao) FindByTotal(condition interface{}) (total int64, err error) {
	var query *gorm.DB

	if conditions, ok := condition.(*QueryOption); ok {
		query = dao.getDBX(conditions)
	} else {
		query = dao.Model(models.Image{}).Where(condition)
	}

	if err = query.Count(&total).Error; err != nil {
		return 0, err
	}
	return
}
