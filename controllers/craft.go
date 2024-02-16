package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/mawarii/sheethappens/models"
)

type CraftInput struct {
	Name string `json:"name" binding:"required"`
}

type CharacterCraftsInput struct {
	Value string `json:"value" binding:"required"`
}

func CreateCraft(c *gin.Context) {
	var input CraftInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newCraft := models.Craft{}

	newCraft.Name = input.Name

	craft, err := newCraft.SaveCraft()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": craft})
}

func GetAllCrafts(c *gin.Context) {
	crafts, err := models.GetAllCraftsByID()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": crafts})
}

func GetCraft(c *gin.Context) {
	craft_id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	craft, err := models.GetCraftByID(uint(craft_id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": craft})
}

func UpdateCraft(c *gin.Context) {
	craft_id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var input CraftInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	craft, err := models.GetCraftByID(uint(craft_id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	craft.Name = input.Name

	new_craft, err := craft.UpdateCraftByID(uint(craft_id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": new_craft})
}

func DeleteCraft(c *gin.Context) {
	craft_id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = models.DeleteCraftByID(uint(craft_id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("craft with id %d deleted", craft_id)})
}
