package models

import (
	"fmt"
	"github.com/miaoming3/wallpaper/global"
	"gorm.io/gorm"
)

type ImageType uint8

const (
	TypeImage ImageType = iota + 1
	TypeVideo ImageType = 2
)

type ImageRecommend uint

const (
	ImageRecommendOff ImageRecommend = iota + 1
	ImageRecommendON  ImageRecommend = 2
)

type Image struct {
	gorm.Model
	Cid         uint
	Category    Category `gorm:"foreignKey:Cid"`
	Name        string
	IsRecommend uint8  `gorm:"is_recommend;type:tinyint(3);default:2;comment:1 推荐 2 不推荐"`
	UserID      uint   // 外键，关联到User表的ID
	User        Users  `gorm:"polymorphic:User;polymorphicValue:User;references:uid` // GORM会自动处理这个关联
	UsersType   uint   `gorm:"user_type;type:tinyint(3);default:1;comment:1 管理员 2 用户"`
	Tags        []Tags `gorm:"many2many:image_tags;joinForeignKey:ImageID;joinReferences:TagsID"`
	Url         string `gorm:"url;type:varchar(255)"`
	Path        string `gorm:"path;type:varchar(255)"`
	Type        uint8  `gorm:"type;type:tinyint(3);default:1;comment:1 图片 2 视频"`
}

func (model *Image) TableName() string {
	return fmt.Sprintf("%s%s", global.SysConfig.Prefix, "image")
}
