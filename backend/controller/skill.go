package controller

// import (
// 	"github.com/gofiber/fiber/v2"
// 	"gitlab.com/Mawarii/sheethappens/database"
// 	"gitlab.com/Mawarii/sheethappens/model"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// )

// func GetSkills(c *fiber.Ctx) error {
// 	coll := database.GetCollection("skills")
// 	cursor, err := coll.Find(c.Context(), bson.D{{}})
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": err.Error(),
// 		})
// 	}

// 	var skills []model.Skill

// 	if err = cursor.All(c.Context(), &skills); err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": err.Error(),
// 		})
// 	}

// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{
// 		"skills": skills,
// 	})
// }

// func GetSkillById(c *fiber.Ctx) error {
// 	id := c.Params("id")
// 	if id == "" {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error": "id is required",
// 		})
// 	}

// 	skillObjectID, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": err.Error(),
// 		})
// 	}

// 	skill := model.Skill{}
// 	coll := database.GetCollection("skills")
// 	err = coll.FindOne(c.Context(), bson.D{{Key: "_id", Value: skillObjectID}}).Decode(&skill)
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": err.Error(),
// 		})
// 	}

// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{
// 		"skill": skill,
// 	})
// }

// type ReqSkill struct {
// 	Name        string `json:"name"                  bson:"name"`
// 	Description string `json:"description,omitempty" bson:"description,omitempty"`
// }

// func CreateSkill(c *fiber.Ctx) error {
// 	b := new(ReqSkill)
// 	if err := c.BodyParser(b); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error": err.Error(),
// 		})
// 	}

// 	coll := database.GetCollection("skills")
// 	result, err := coll.InsertOne(c.Context(), b)
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"message": "failed to create skill",
// 			"error":   err.Error(),
// 		})
// 	}

// 	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
// 		"result": result,
// 	})
// }

// func UpdateSkill(c *fiber.Ctx) error {
// 	b := new(ReqSkill)
// 	if err := c.BodyParser(b); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error": err.Error(),
// 		})
// 	}

// 	id := c.Params("id")
// 	if id == "" {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error": "id is required",
// 		})
// 	}

// 	skillObjectID, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": err.Error(),
// 		})
// 	}

// 	coll := database.GetCollection("skills")
// 	result, err := coll.UpdateOne(c.Context(), bson.D{{Key: "_id", Value: skillObjectID}}, bson.D{{Key: "$set", Value: b}})
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"message": "failed to updated skill",
// 			"error":   err.Error(),
// 		})
// 	}

// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{
// 		"result": result,
// 	})
// }

// func DeleteSkill(c *fiber.Ctx) error {
// 	id := c.Params("id")
// 	if id == "" {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error": "id is required",
// 		})
// 	}

// 	skillObjectID, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": err.Error(),
// 		})
// 	}

// 	coll := database.GetCollection("skills")
// 	result, err := coll.DeleteOne(c.Context(), bson.D{{Key: "_id", Value: skillObjectID}})
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"message": "failed to delete skill",
// 			"error":   err.Error(),
// 		})
// 	}

// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{
// 		"result": result,
// 	})
// }
