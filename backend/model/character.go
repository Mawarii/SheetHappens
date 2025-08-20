package model

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Character struct {
	gorm.Model
	UserID   uint           `gorm:"not null;index" json:"user_id"`
	SystemID uint           `gorm:"not null;index" json:"system_id"`
	Name     string         `gorm:"not null;" json:"name"`
	Data     datatypes.JSON `gorm:"not null;type:jsonb" json:"data"`
	System   System         `gorm:"foreignKey:SystemID" json:"system"`
}
