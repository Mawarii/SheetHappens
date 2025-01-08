package model

import (
	"gorm.io/gorm"
)

type Skill struct {
	gorm.Model
	Name        string          `gorm:"not null;unique;" json:"name"`
	Description string          `json:"description,omitempty"`
	Categories  []SkillCategory `gorm:"many2many:skills_skillcategories;" json:"categories"`
}

type SkillCategory struct {
	gorm.Model
	Name        string  `gorm:"not null;unique;" json:"name"`
	Description string  `json:"description,omitempty"`
	Skills      []Skill `gorm:"many2many:skills_skillcategories;" json:"skills"`
}

type CharacterSkill struct {
	gorm.Model
	CharacterID uint `json:"character_id"`
	SkillID     uint `json:"skill_id"`
	Level       uint `gorm:"not null;" json:"level"`
}
