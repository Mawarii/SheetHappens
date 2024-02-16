package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/mawarii/sheethappens/models"
	"gitlab.com/mawarii/sheethappens/utils/token"
)

type CharacterInput struct {
	Name         string `json:"name" binding:"required"`
	Level        uint   `json:"level"`
	Health       int    `json:"health" binding:"required"`
	MentalHealth int    `json:"mental health" binding:"required"`
	Race         string `json:"race" binding:"required"`
	Gender       string `json:"gender"`
	Height       string `json:"height" binding:"required"`
	Weight       string `json:"weight" binding:"required"`
	Dodge        uint   `json:"dodge" binding:"required"`
}

func CreateCharacter(c *gin.Context) {
	user_id, err := token.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var input CharacterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	char := models.Character{}

	char.Name = input.Name
	char.Level = input.Level
	char.Health = input.Health
	char.MentalHealth = input.MentalHealth
	char.Race = input.Race
	char.Gender = input.Gender
	char.Height = input.Height
	char.Weight = input.Weight
	char.Dodge = input.Dodge
	char.UserID = user_id

	character, err := char.SaveCharacter()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": character})
}

func GetAllCharacters(c *gin.Context) {
	user_id, err := token.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	characters, err := models.GetAllCharactersByUserID(user_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": characters})
}

func GetCharacter(c *gin.Context) {
	user_id, err := token.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	char_id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	char, err := models.GetCharacterByID(user_id, uint(char_id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": char})
}

func UpdateCharacter(c *gin.Context) {
	user_id, err := token.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	char_id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var input CharacterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	char, err := models.GetCharacterByID(user_id, uint(char_id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	char.Name = input.Name
	char.Level = input.Level
	char.Health = input.Health
	char.MentalHealth = input.MentalHealth
	char.Race = input.Race
	char.Gender = input.Gender
	char.Height = input.Height
	char.Weight = input.Weight
	char.Dodge = input.Dodge
	char.UserID = user_id

	character, err := char.UpdateCharacterByID(user_id, uint(char_id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": character})
}

func DeleteCharacter(c *gin.Context) {
	user_id, err := token.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	char_id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = models.DeleteCharacterByID(user_id, uint(char_id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("character with id %d deleted", char_id)})
}
