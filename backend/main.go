package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"gitlab.com/Mawarii/sheethappens/controller"
	"gitlab.com/Mawarii/sheethappens/database"
	"gitlab.com/Mawarii/sheethappens/middleware"
	"gitlab.com/Mawarii/sheethappens/utils"
)

func main() {
	utils.LoadEnv()
	database.Init()
	defer database.CloseConnection()

	app := fiber.New()

	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:8080",
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "Content-Type, Authorization",
	}))

	api := app.Group("/api")

	auth := api.Group("/auth")
	auth.Post("/register", controller.Register)
	auth.Post("/login", controller.Login)

	char := api.Group("/characters", middleware.JWTProtected)
	char.Get("/", controller.GetCharacters)
	char.Get("/:id", controller.GetCharacterById)
	char.Post("/", controller.CreateCharacter)
	char.Put("/:id", controller.UpdateCharacter)
	char.Delete("/:id", controller.DeleteCharacter)

	app.Listen(":3000")
}
