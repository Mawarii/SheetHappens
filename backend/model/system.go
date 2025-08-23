package model

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type System struct {
	gorm.Model
	Name        string         `gorm:"not null;unique" json:"name"`
	DisplayName string         `gorm:"not null;unique;" json:"display_name"`
	Schema      datatypes.JSON `gorm:"type:jsonb;not null" json:"schema"`
	Description string         `json:"description,omitempty"`
}
