package models

import (
	"gorm.io/gorm"
)

type Photo struct {
	gorm.Model
	Title    string `gorm:"not null" json:"title" form:"title"`
	Caption  string `json:"caption" form:"caption"`
	PhotoURL string `gorm:"not null" json:"photo_url"`
	UserID   uint   `gorm:"not null" json:"-"`
	User     User   `json:"-"`
}
