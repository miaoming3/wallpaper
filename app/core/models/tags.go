package models

import (
	"fmt"
	"github.com/miaoming3/wallpaper/app/global"
)

type Tags struct {
	ID     uint   `gorm:"primaryKey"`
	Name   string `gorm:"uniqueIndex;type:varchar(255)"` // 假设标签名是唯一的
	UserID uint   // 外键，关联到User表的ID
	User   Users  `gorm:"references:uid"` // GORM会自动处理这个关联
}

func (model *Tags) TableName() string {
	return fmt.Sprintf("%s%s", global.SysConfig.Prefix, "tags")
}

type ImageTags struct {
	ImageID uint `gorm:"primaryKey;column:image_id"`
	TagsID  uint `gorm:"primaryKey;column:tags_id"`
	// 通常不需要其他字段，但如果有额外信息可以添加
}

func (model *ImageTags) TableName() string {
	return fmt.Sprintf("%s%s", global.SysConfig.Prefix, "image_tags")
}
