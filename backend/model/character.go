package model

import (
	"gorm.io/gorm"
)

type Character struct {
	gorm.Model
	Name         string           `gorm:"not null;" json:"name"`
	Level        uint             `json:"level,omitempty"`
	Health       int              `json:"health,omitempty"`
	MentalHealth int              `json:"mental_health,omitempty"`
	Mana         uint             `json:"mana,omitempty"`
	Race         string           `json:"race,omitempty"`
	Gender       string           `json:"gender,omitempty"`
	Height       string           `json:"height,omitempty"`
	Weight       string           `json:"weight,omitempty"`
	Dodge        uint             `json:"dodge,omitempty"`
	UserID       uint             `json:"user_id"`
	Skills       []CharacterSkill `gorm:"foreignKey:CharacterID" json:"skills,omitempty"`
}
