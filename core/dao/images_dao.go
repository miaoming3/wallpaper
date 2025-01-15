package dao

import (
	"github.com/miaoming3/wallpaper/controller/dto"
	models2 "github.com/miaoming3/wallpaper/core/models"
	"github.com/miaoming3/wallpaper/http/global"
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
	query := dao.Model(models2.Image{})
	for key, value := range condition.Conditions {
		if key == "tags_id" && value != nil {
			subQuery := dao.Model(models2.ImageTags{}).Select("image_id").Where("tags_id = ?", value)
			query = query.Where("id IN (?)", subQuery)
		} else if key == "username" && value != nil {
			subQuery := dao.Model(models2.Users{}).Select("uid").Where("username like ?", value)
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

func (dao *ImagesDao) FindByAll(condition *QueryOption, page int, pageSize int) ([]models2.Image, error) {
	var images []models2.Image
	offset := (page - 1) * pageSize

	if err := dao.getDBX(condition).Offset(offset).Limit(pageSize).Find(&images).Error; err != nil {
		return nil, err
	}
	return images, nil
}
func (dao *ImagesDao) SaveImage(data *dto.ImageSave, uid uint) error {
	image := &models2.Image{
		Name:        data.Name,
		Type:        data.Type,
		Cid:         data.CID,
		Url:         data.Url,
		Path:        data.Path,
		IsRecommend: data.IsRecommend,
		UserID:      uid,
	}
	return dao.Begin().Transaction(func(db *gorm.DB) error {
		if err := db.Model(models2.Image{}).Create(image).Error; err != nil {
			db.Rollback()
			return err
		}
		if len(data.Tags) > 0 {
			var tags []models2.Tags
			if err := db.Model(models2.Image{}).Where("id IN ?", data.Tags).Find(&tags).Error; err != nil {
				db.Rollback()
				return err
			}

			if err := db.Model(models2.Image{}).Model(image).Association("Tags").Append(tags); err != nil {
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
	image := &models2.Image{
		Name:        data.Name,
		Type:        data.Type,
		Cid:         data.CID,
		Url:         data.Url,
		Path:        data.Path,
		IsRecommend: data.IsRecommend,
	}
	return dao.Begin().Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(models2.Image{}).Where("id = ?", data.ID).Updates(image).Error; err != nil {
			tx.Rollback()
			return err
		}
		if len(data.Tags) > 0 {
			var tags []models2.Tags
			if err := tx.Model(models2.Image{}).Where("id IN ?", data.Tags).Find(&tags).Error; err != nil {
				tx.Rollback()
				return err
			}
			// 使用 Replace 关联查询出来的 Tags 实例
			if err := tx.Model(&models2.Image{Model: gorm.Model{ID: data.ID}}).Association("Tags").Replace(tags); err != nil {
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
		query = dao.Model(models2.Image{}).Where(condition)
	}

	if err = query.Count(&total).Error; err != nil {
		return 0, err
	}
	return
}
