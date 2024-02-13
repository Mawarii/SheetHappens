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

	protected := router.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", controllers.CurrentUser)

	router.Run(":8080")
}
