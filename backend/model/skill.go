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
	CharacterID     uint `gorm:"uniqueIndex:idx_character_skill_category" json:"character_id"`
	SkillID         uint `gorm:"uniqueIndex:idx_character_skill_category" json:"skill_id"`
	SkillCategoryID uint `gorm:"uniqueIndex:idx_character_skill_category" json:"skill_category_id"`
	Level           uint `gorm:"not null;"                                json:"level"`
}
