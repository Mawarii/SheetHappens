package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gitlab.com/Mawarii/sheethappens/database"
	"gitlab.com/Mawarii/sheethappens/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetCharacters(c *fiber.Ctx) error {
	userToken := c.Locals("jwt").(*jwt.Token)
	claims := userToken.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(string)

	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	coll := database.GetCollection("characters")
	cursor, err := coll.Find(c.Context(), bson.D{{Key: "user_id", Value: userObjectID}})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var characters []model.Character

	if err = cursor.All(c.Context(), &characters); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"characters": characters,
	})
}

func GetCharacterById(c *fiber.Ctx) error {
	userToken := c.Locals("jwt").(*jwt.Token)
	claims := userToken.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(string)

	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "id is required",
		})
	}

	charObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	character := model.Character{}
	coll := database.GetCollection("characters")
	err = coll.FindOne(c.Context(), bson.D{{Key: "user_id", Value: userObjectID}, {Key: "_id", Value: charObjectID}}).Decode(&character)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"character": character,
	})
}

type ReqCharacter struct {
	UserID       primitive.ObjectID                `json:"user_id"                bson:"user_id"`
	Name         string                            `json:"name"                   bson:"name"`
	Level        int                               `json:"level,omitempty"        bson:"level,omitempty"`
	Health       int                               `json:"health,omitempty"       bson:"health,omitempty"`
	MentalHealth int                               `json:"mentalhealth,omitempty" bson:"mentalhealth,omitempty"`
	Mana         int                               `json:"mana,omitempty"         bson:"mana,omitempty"`
	Race         string                            `json:"race,omitempty"         bson:"race,omitempty"`
	Gender       string                            `json:"gender,omitempty"       bson:"gender,omitempty"`
	Height       string                            `json:"height,omitempty"       bson:"height,omitempty"`
	Weight       string                            `json:"weight,omitempty"       bson:"weight,omitempty"`
	Dodge        int                               `json:"dodge,omitempty"        bson:"dodge,omitempty"`
	Skills       map[string][]model.SkillValuePair `json:"skills,omitempty"       bson:"skills,omitempty"`
}

func CreateCharacter(c *fiber.Ctx) error {
	userToken := c.Locals("jwt").(*jwt.Token)
	claims := userToken.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(string)

	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	b := new(ReqCharacter)
	if err := c.BodyParser(b); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	b.UserID = userObjectID

	coll := database.GetCollection("characters")
	result, err := coll.InsertOne(c.Context(), b)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to create character",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"result": result,
	})
}

func UpdateCharacter(c *fiber.Ctx) error {
	userToken := c.Locals("jwt").(*jwt.Token)
	claims := userToken.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(string)

	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	b := new(ReqCharacter)
	if err := c.BodyParser(b); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	b.UserID = userObjectID

	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "id is required",
		})
	}

	charObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	coll := database.GetCollection("characters")
	result, err := coll.UpdateOne(c.Context(), bson.D{{Key: "user_id", Value: userObjectID}, {Key: "_id", Value: charObjectID}}, bson.D{{Key: "$set", Value: b}})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to updated character",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"result": result,
	})
}

func DeleteCharacter(c *fiber.Ctx) error {
	userToken := c.Locals("jwt").(*jwt.Token)
	claims := userToken.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(string)

	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "id is required",
		})
	}

	charObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	coll := database.GetCollection("characters")
	result, err := coll.DeleteOne(c.Context(), bson.D{{Key: "user_id", Value: userObjectID}, {Key: "_id", Value: charObjectID}})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to delete character",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"result": result,
	})
}
