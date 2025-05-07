package controller

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gitlab.com/Mawarii/sheethappens/database"
	"gitlab.com/Mawarii/sheethappens/model"
	"gorm.io/datatypes"
)

func GetCharacters(c *fiber.Ctx) error {
	userToken := c.Locals("jwt").(*jwt.Token)
	claims := userToken.Claims.(jwt.MapClaims)
	userID, ok := claims["user_id"]

	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid token",
		})
	}

	var characters []model.Character

	result := database.DB().Where("user_id = ?", userID).Find(&characters)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to get characters",
			"details": result.Error.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(characters)
}

func GetCharacterById(c *fiber.Ctx) error {
	userToken := c.Locals("jwt").(*jwt.Token)
	claims := userToken.Claims.(jwt.MapClaims)
	userID, ok := claims["user_id"]

	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid token",
		})
	}

	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "id is required",
		})
	}

	var character model.Character

	result := database.DB().Model(model.Character{}).Where("id = ? AND user_id = ?", id, userID).First(&character)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to get character",
			"details": result.Error.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(character)
}

type ReqCharacter struct {
	Name   string         `gorm:"not null;" json:"name"`
	System string         `json:"system"`
	Data   datatypes.JSON `gorm:"type:jsonb" json:"data"`
}

func CreateCharacter(c *fiber.Ctx) error {
	userToken := c.Locals("jwt").(*jwt.Token)
	claims := userToken.Claims.(jwt.MapClaims)
	userID, ok := claims["user_id"].(float64)

	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid token",
		})
	}

	body := new(ReqCharacter)

	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if !json.Valid(body.Data) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid JSON in data",
		})
	}

	var character model.Character
	character.UserID = uint(userID)
	character.Name = body.Name
	character.System = body.System
	character.Data = body.Data

	result := database.DB().Create(&character)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to create character",
			"details": result.Error.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"character": character,
	})
}

func UpdateCharacter(c *fiber.Ctx) error {
	userToken := c.Locals("jwt").(*jwt.Token)
	claims := userToken.Claims.(jwt.MapClaims)
	userID, ok := claims["user_id"]

	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid token",
		})
	}

	body := new(ReqCharacter)

	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if !json.Valid(body.Data) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid JSON in data",
		})
	}

	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "id is required",
		})
	}

	var character model.Character

	result := database.DB().Model(model.Character{}).Where("id = ? AND user_id = ?", id, userID).First(&character)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to update character",
			"details": result.Error.Error(),
		})
	}

	character.Name = body.Name
	character.System = body.System
	character.Data = body.Data

	database.DB().Save(&character)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"character": character,
	})
}

func DeleteCharacter(c *fiber.Ctx) error {
	userToken := c.Locals("jwt").(*jwt.Token)
	claims := userToken.Claims.(jwt.MapClaims)
	userID, ok := claims["user_id"]

	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid token",
		})
	}

	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "id is required",
		})
	}

	result := database.DB().Unscoped().Where("id = ? AND user_id = ?", id, userID).Delete(&model.Character{})
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to delete character",
			"details": result.Error.Error(),
		})
	}

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Character already deleted",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": "Character deleted successfully",
	})
}
