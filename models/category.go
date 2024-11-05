package models

import (
	"encoding/json"
	"fmt"
	"github.com/miaoming3/wallpaper/global"
	"gorm.io/gorm"
)

type CategoryStatus uint

const (
	StatusOn  CategoryStatus = iota + 1
	StatusOff CategoryStatus = 2
)

type Category struct {
	gorm.Model
	Name   string `gorm:"name;type:varchar(255) "`
	Pid    uint
	Sort   uint
	Status uint `gorm:"status;type:tinyint(3);default:1"`
}

func (ca *Category) TableName() string {
	return fmt.Sprintf("%s%s", global.SysConfig.Prefix, "catgegory")
}

func (s *CategoryStatus) UnmarshalJSON(b []byte) error {
	var status uint
	if err := json.Unmarshal(b, &status); err != nil {
		return err
	}
	switch status {
	case 1:
		*s = StatusOn
	case 2:
		*s = StatusOn
	default:
		*s = StatusOn
	}
	return nil
}
