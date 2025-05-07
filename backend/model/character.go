package model

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Character struct {
	gorm.Model
	UserID uint           `json:"user_id"`
	Name   string         `gorm:"not null;" json:"name"`
	System string         `json:"system"`
	Data   datatypes.JSON `gorm:"type:jsonb" json:"data"`
}
