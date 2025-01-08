package model

import (
	"gorm.io/gorm"
)

type Skill struct {
	gorm.Model
	Name        string `gorm:"not null;unique;" json:"name"`
	Description string `json:"description,omitempty"`
}

type SkillCategory struct {
	gorm.Model
	Name        string `gorm:"not null;unique;" json:"name"`
	Description string `json:"description,omitempty"`
}

type CharacterSkillCategory struct {
	gorm.Model
	CharacterID     uint          `json:"character_id"`
	SkillID         uint          `json:"skill_id"`
	SkillCategoryID uint          `json:"skill_category_id"`
	Level           uint          `gorm:"not null;" json:"level"`
	Character       Character     `json:"character,omitempty"`
	Skill           Skill         `json:"skill,omitempty"`
	SkillCategory   SkillCategory `json:"skill_category,omitempty"`
}
