package dao

import (
	"github.com/miaoming3/wallpaper/app/global"
	"github.com/miaoming3/wallpaper/app/models"
	"gorm.io/gorm"
	"strings"
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

func (dao *AdminDao) getQuery(condition map[string]interface{}) (query *gorm.DB) {
	query = dao.Model(&models.AdminModel{})
	if len(condition) > 0 {
		for k, v := range condition {
			if strings.HasPrefix(k, "or") {
				k = strings.ReplaceAll(k, "or", "")
				query.Or(k, v)
			} else {
				query.Where(k, v)
			}
		}
	}
	return
}

func (dao *AdminDao) FindById(condition map[string]interface{}) (admin *models.AdminModel, err error) {

	if err = dao.getQuery(condition).First(&admin).Error; err != nil {
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

func (dao *AdminDao) GetList(condition map[string]interface{}, page int, pageSize int) (adminModel []*models.AdminModel, err error) {
	if err = dao.getQuery(condition).Offset(page - 1).Limit(pageSize).Find(&adminModel).Error; err != nil {
		return nil, err
	}
	return
}
func (dao *AdminDao) GetTotal(condition map[string]interface{}) (total int64, err error) {
	if err = dao.getQuery(condition).Count(&total).Error; err != nil {
		return 0, err
	}
	return
}

func (dao *AdminDao) DeleteRow(id int, forever bool, adminModel *models.AdminModel) error {
	query := dao.Model(&models.AdminModel{})
	if forever {
		query.Unscoped()
	}
	return query.Delete(&adminModel, id).Error

}
