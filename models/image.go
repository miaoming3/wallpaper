package models

import (
	"fmt"
	"github.com/miaoming3/wallpaper/global"
	"gorm.io/gorm"
)

type Image struct {
	gorm.Model
	Cid         uint
	Category    Category `gorm:"foreignKey:Cid"`
	Name        string
	IsRecommend uint
	UserID      uint   // 外键，关联到User表的ID
	User        Users  `gorm:"foreignKey:UserID"` // GORM会自动处理这个关联
	Tags        []Tags `gorm:"many2many:image_tags;"`
}

func (model *Image) TableName() string {
	return fmt.Sprintf("%s%s", global.SysConfig.Prefix, "image")
}
