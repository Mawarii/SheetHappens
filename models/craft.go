package models

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Craft struct {
	gorm.Model
	Name string `gorm:"size:255;not null;unique;" json:"name"`
}

type CharacterCraft struct {
	CharacterID int `gorm:"primaryKey"`
	CraftID     int `gorm:"primaryKey"`
	CreatedAt   time.Time
	DeletedAt   gorm.DeletedAt
	Value       int `gorm:"not null;unique" json:"value"`
}

func (craft *Craft) SaveCraft() (*Craft, error) {
	newCraftName := craft.Name

	result := database.Where("LOWER(name) = ?", strings.ToLower(craft.Name)).FirstOrCreate(&craft)

	if result.Error != nil {
		return &Craft{}, result.Error
	}

	if result.RowsAffected == 0 {
		return craft, fmt.Errorf("craft '%s' already exists", newCraftName)
	}

	return craft, nil
}

func GetAllCraftsByID() ([]Craft, error) {
	var crafts []Craft

	if err := database.Model(Craft{}).Find(&crafts).Error; err != nil {
		return crafts, errors.New("no crafts found")
	}

	return crafts, nil
}

func GetCraftByID(craft_id uint) (Craft, error) {
	var craft Craft

	if err := database.Model(Craft{}).First(&craft, craft_id).Error; err != nil {
		return craft, errors.New("craft not found")
	}

	return craft, nil
}

func (craft *Craft) UpdateCraftByID(craft_id uint) (*Craft, error) {

	if err := database.Save(&craft).Error; err != nil {
		return craft, errors.New("craft not found")
	}

	return craft, nil
}

func DeleteCraftByID(craft_id uint) error {
	var craft Craft

	result := database.Where("id = ?", craft_id).Delete(&craft)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("craft not found")
	}

	return nil
}
