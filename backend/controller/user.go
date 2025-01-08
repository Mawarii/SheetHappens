package controller

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gitlab.com/Mawarii/sheethappens/database"
	"gitlab.com/Mawarii/sheethappens/model"
	"gitlab.com/Mawarii/sheethappens/utils"
)

type authRequest struct {
	Username string `gorm:"not null;unique;"  json:"username"`
	Password string `gorm:"not null;"         json:"password"`
}

func Register(c *fiber.Ctx) error {
	body := new(authRequest)
	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if body.Username == "" || body.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Username and Password required!",
		})
	}

	body.Username = strings.ToLower(body.Username)

	var user = model.User{
		Username: body.Username,
		Password: utils.GeneratePassword(body.Password),
	}

	result := database.DB().Create(&user)

	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": result.Error,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"user_id": user.ID,
	})
}

func Login(c *fiber.Ctx) error {
	body := new(authRequest)
	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	user := model.User{}
	result := database.DB().Where("username = LOWER(?)", body.Username).First(&user)
	if result.Error != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid credentials",
		})
	}

	if !utils.ComparePassword(user.Password, body.Password) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid credentials",
		})
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "access token generation failed",
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    token,
		HTTPOnly: true,
		Secure:   true,
		Expires:  time.Now().Add(24 * time.Hour),
	})

	return c.JSON(fiber.Map{
		"message": "login successful",
	})
}

func Logout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    "",
		HTTPOnly: true,
		Secure:   true,
		Expires:  time.Now().Add(-time.Hour),
	})

	return c.JSON(fiber.Map{
		"message": "logout successful",
	})
}

func GetUserInfo(c *fiber.Ctx) error {
	userToken := c.Locals("jwt").(*jwt.Token)
	claims := userToken.Claims.(jwt.MapClaims)
	userID := uint(claims["user_id"].(float64))

	var user model.User

	result := database.DB().Preload("Characters").First(&user, userID)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": result.Error,
		})
	}

	user.Password = "REDACTED"

	return c.JSON(user)
}
