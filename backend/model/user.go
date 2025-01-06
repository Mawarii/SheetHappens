package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username   string      `gorm:"not null;unique;"  json:"username"`
	Password   string      `gorm:"not null;"         json:"password"`
	Characters []Character `gorm:"foreignKey:UserID" json:"characters,omitempty"`
}
