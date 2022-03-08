package main

import (
	"Minerva/database"
	"Minerva/routesGlobal"
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
		port = "3000"
	}
	errListen := app.Listen(":3000")
	if errListen != nil {
		os.Exit(1)
	}

}
