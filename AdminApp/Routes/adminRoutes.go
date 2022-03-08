package Routes

import (
	adminHandler "Minerva/AdminApp/handler"
	"Minerva/config"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func SetupNoteRoutes(router fiber.Router) {
	admin := router.Group("/admin")

	// JWT Middleware
	admin.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(config.Config("SECRET")),
	}))
	// Create Category
	admin.Post("/new-category", adminHandler.CreateCategory)
	// All Category
	admin.Get("/all-category", adminHandler.GetAllCategory)
}
