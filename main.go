package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gitlab.com/mawarii/sheethappens/controllers"
	"gitlab.com/mawarii/sheethappens/middlewares"
	"gitlab.com/mawarii/sheethappens/models"
)

func main() {
	models.DatabaseInit()

	router := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	public := router.Group("/api")

	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	protected := router.Group("/api/user")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/info", controllers.CurrentUser)
	protected.POST("/characters", controllers.CreateCharacter)
	protected.GET("/characters", controllers.AllCharacters)

	router.Run(":8080")
}
