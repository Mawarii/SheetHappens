package models

import (
	"errors"

	"gorm.io/gorm"
)

type Character struct {
	gorm.Model
	Name         string `gorm:"size:255;not null;" json:"name"`
	Level        uint   `json:"level"`
	Health       int    `gorm:"not null;" json:"health"`
	MentalHealth int    `gorm:"not null;" json:"mental health"`
	Race         string `gorm:"size:255;not null;" json:"race"`
	Gender       string `gorm:"size:255;" json:"gender"`
	Height       string `gorm:"size:255;not null;" json:"height"`
	Weight       string `gorm:"size:255;not null;" json:"weight"`
	Dodge        uint   `gorm:"not null;" json:"dodge"`
	UserID       uint
	Crafts       []Craft `gorm:"many2many:character_crafts;"`
}

func (char *Character) SaveCharacter() (*Character, error) {
	var err error = database.Create(&char).Error

	if err != nil {
		return &Character{}, err
	}

	return char, nil
}

func GetAllCharactersByUserID(uid uint) ([]Character, error) {
	var chars []Character

	if err := database.Model(Character{}).Where("user_id = ?", uid).Find(&chars).Error; err != nil {
		return chars, errors.New("characters not found")
	}

	return chars, nil
}

func GetCharacterByID(user_id uint, char_id uint) (Character, error) {
	var char Character

	if err := database.Model(Character{}).Where("user_id = ?", user_id).First(&char, char_id).Error; err != nil {
		return char, errors.New("character not found or not owned by this user")
	}

	return char, nil
}

func (char *Character) UpdateCharacterByID(user_id uint, char_id uint) (*Character, error) {

	if err := database.Save(&char).Error; err != nil {
		return char, errors.New("character not found or not owned by this user")
	}

	return char, nil
}

func DeleteCharacterByID(user_id uint, char_id uint) error {
	var char Character

	result := database.Where("user_id = ? AND id = ?", user_id, char_id).Delete(&char)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("character not found or not owned by this user")
	}

	return nil
}
