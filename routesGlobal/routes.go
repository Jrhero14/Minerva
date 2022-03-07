package routesGlobal

import (
	authRoutes "Minerva/AuthApp/routes"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	// Setup the Node Routes
	authRoutes.SetupNoteRoutes(api)
}
