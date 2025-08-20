package controller

import (
	"encoding/json"
	"strings"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/Mawarii/sheethappens/database"
	"gitlab.com/Mawarii/sheethappens/model"
	"gorm.io/datatypes"
)

func GetSystems(c *fiber.Ctx) error {
	var systems []model.System

	result := database.DB().Find(&systems)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to get systems",
			"details": result.Error.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(systems)
}

func GetSystemById(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "id is required",
		})
	}

	var system model.System

	result := database.DB().Model(model.System{}).Where("id = ?", id).First(&system)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to get system",
			"details": result.Error.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(system)
}

type ReqSystem struct {
	Name        string         `json:"name"`
	Schema      datatypes.JSON `json:"schema"`
	Description string         `json:"description"`
}

func CreateSystem(c *fiber.Ctx) error {
	body := new(ReqSystem)

	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Failed to create system",
			"details": err.Error(),
		})
	}

	if !json.Valid(body.Schema) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid JSON in data",
		})
	}

	var system model.System
	system.Name = strings.ToLower(body.Name)
	system.DisplayName = body.Name
	system.Schema = body.Schema
	system.Description = body.Description

	result := database.DB().Create(&system)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to create system",
			"details": result.Error.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"system": system,
	})
}

func UpdateSystem(c *fiber.Ctx) error {
	body := new(ReqSystem)

	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Failed to update system",
			"details": err.Error(),
		})
	}

	if !json.Valid(body.Schema) {
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

	var system model.System

	result := database.DB().Model(model.System{}).Where("id = ?", id).First(&system)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to update system",
			"details": result.Error.Error(),
		})
	}

	system.Name = strings.ToLower(body.Name)
	system.DisplayName = body.Name
	system.Schema = body.Schema
	system.Description = body.Description

	database.DB().Save(&system)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"system": system,
	})
}

func DeleteSystem(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "id is required",
		})
	}

	// If you want to disable soft deletion use: database.DB().Unscoped().Where(...)
	result := database.DB().Where("id = ?", id).Delete(&model.System{})
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to delete system",
			"details": result.Error.Error(),
		})
	}

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "System already deleted",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": "System deleted successfully",
	})
}
