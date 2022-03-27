package main

import (
	"github.com/Jrhero14/Minerva/database"
	"github.com/Jrhero14/Minerva/routesGlobal"
	"github.com/gofiber/fiber/v2"
	"os"
)

func main() {
	// Start a new fiber app
	app := fiber.New()

	// Connect to the Database
	database.ConnectDB()

	// Setup the router
	routesGlobal.SetupRoutes(app)

	// Listen on PORT 3000
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	errListen := app.Listen(":8000")
	if errListen != nil {
		os.Exit(1)
	}

}
