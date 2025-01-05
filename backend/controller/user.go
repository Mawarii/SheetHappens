package controller

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/Mawarii/sheethappens/database"
	"gitlab.com/Mawarii/sheethappens/model"
	"gitlab.com/Mawarii/sheethappens/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type authRequest struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

func Register(c *fiber.Ctx) error {
	b := new(authRequest)
	if err := c.BodyParser(b); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if b.Username == "" || b.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Username and Password required!",
		})
	}

	var user = model.User{
		ID:       primitive.NewObjectID(),
		Username: b.Username,
		Password: utils.GeneratePassword(b.Password),
	}

	coll := database.GetCollection("users")
	result, err := coll.InsertOne(c.Context(), user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to create user",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"result": result,
	})
}

func Login(c *fiber.Ctx) error {
	b := new(authRequest)
	if err := c.BodyParser(b); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	user := model.User{}
	coll := database.GetCollection("users")
	err := coll.FindOne(c.Context(), bson.M{"username": b.Username}).Decode(&user)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid credentials",
		})
	}

	if !utils.ComparePassword(user.Password, b.Password) {
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
