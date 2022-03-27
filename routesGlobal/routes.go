package routesGlobal

import (
	adminRoutes "github.com/Jrhero14/Minerva/AdminApp/Routes"
	authRoutes "github.com/Jrhero14/Minerva/AuthApp/routes"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	// Setup the Node Routes
	authRoutes.SetupNoteRoutes(api)
	adminRoutes.SetupNoteRoutes(api)
}
