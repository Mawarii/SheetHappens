package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username    string      `gorm:"not null;unique;" json:"username"`
	DisplayName string      `gorm:"not null;unique;" json:"display_name"`
	Password    string      `gorm:"not null;" json:"-"`
	Characters  []Character `gorm:"foreignKey:UserID;" json:"characters,omitempty"`
}
