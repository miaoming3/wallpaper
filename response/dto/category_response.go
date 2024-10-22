package dto

import "github.com/miaoming3/wallpaper/models"

type CategoryListResponse struct {
	Name    string                  `json:"name"`
	ID      uint                    `json:"id"`
	Pid     uint                    `json:"pid"`
	Status  models.CategoryStatus   `json:"status"`
	Sort    uint                    `json:"sort"`
	Childer []*CategoryListResponse `json:"childer,omitempty"`
}
