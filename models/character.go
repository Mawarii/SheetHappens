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
}

func GetAllCharacters(uid uint) ([]Character, error) {
	var chars []Character

	if err := database.Model(Character{}).Where("user_id = ?", uid).Find(&chars).Error; err != nil {
		return chars, errors.New("character not found")
	}

	return chars, nil
}

func (char *Character) SaveCharacter() (*Character, error) {
	var err error = database.Create(&char).Error

	if err != nil {
		return &Character{}, err
	}

	return char, nil
}
