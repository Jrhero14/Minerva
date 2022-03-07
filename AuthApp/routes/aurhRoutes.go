package routes

import (
	authHandler "Minerva/AuthApp/handler"
	"Minerva/config"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func SetupNoteRoutes(router fiber.Router) {
	auth := router.Group("/auth")

	// login
	auth.Post("/login", authHandler.Login)
	// Create a User
	auth.Post("/regis", authHandler.CreateUser)
	// JWT Middleware
	auth.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(config.Config("SECRET")),
	}))
	// Update Member
	auth.Post("/update-member-data", authHandler.UpdateMember)
	// Get All User
	auth.Get("/get-all-user", authHandler.AllUser)
	// Get Info Member User
	auth.Post("/get-info-member", authHandler.GetinfoMember)
}
