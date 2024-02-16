package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/mawarii/sheethappens/controllers"
	"gitlab.com/mawarii/sheethappens/middlewares"
	"gitlab.com/mawarii/sheethappens/models"
)

func main() {
	models.DatabaseInit()

	router := gin.Default()

	public := router.Group("/api")

	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	protected := router.Group("/api/user")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/info", controllers.CurrentUser)

	protected.GET("/characters", controllers.GetAllCharacters)

	protected.POST("/character", controllers.CreateCharacter)
	protected.GET("/character/:id", controllers.GetCharacter)
	protected.PUT("/character/:id", controllers.UpdateCharacter)
	protected.DELETE("/character/:id", controllers.DeleteCharacter)

	protected.GET("/crafts", controllers.GetAllCrafts)

	protected.POST("/craft", controllers.CreateCraft)
	protected.GET("/craft/:id", controllers.GetCraft)
	protected.PUT("/craft/:id", controllers.UpdateCraft)
	protected.DELETE("/craft/:id", controllers.DeleteCraft)

	router.Run(":8080")
}
