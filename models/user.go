package models

import (
	"errors"
	"fmt"
	"html"
	"strings"

	"gitlab.com/mawarii/sheethappens/utils/token"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username   string `gorm:"size:255;not null;unique" json:"username"`
	Password   string `gorm:"size:255;not null;" json:"password"`
	Characters []Character
}

func GetUserByID(uid uint) (User, error) {
	var user User

	if err := database.Preload("Characters").First(&user, uid).Error; err != nil {
		return user, errors.New("user not found")
	}

	user.PrepareGive()

	return user, nil
}

func (user *User) PrepareGive() {
	user.Password = ""
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(username string, password string) (string, error) {
	var err error

	user := User{}

	err = database.Model(User{}).Where("username = ?", username).Take(&user).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, user.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := token.GenerateToken(user.ID)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (user *User) SaveUser() (*User, error) {
	newUsername := user.Username

	result := database.Where("LOWER(username) = ?", strings.ToLower(user.Username)).FirstOrCreate(&user)

	if result.Error != nil {
		return &User{}, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("username '%s' already in use", newUsername)
	}

	return user, nil
}

func (user *User) BeforeSave(tx *gorm.DB) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))

	return nil
}
