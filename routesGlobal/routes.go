package routesGlobal

import (
	adminRoutes "github.com/Jrhero14/Minerva/AdminApp/Routes"
	authRoutes "github.com/Jrhero14/Minerva/AuthApp/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	// Setup the Node Routes
	// Or extend your config for customization
	api.Use(cors.New(cors.Config{
		Next:             nil,
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:     "",
		AllowCredentials: true,
		ExposeHeaders:    "",
		MaxAge:           0,
	}))
	authRoutes.SetupNoteRoutes(api)
	adminRoutes.SetupNoteRoutes(api)
}
