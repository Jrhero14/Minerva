package main

import (
	"Minerva/database"
	"Minerva/routesGlobal"
	"github.com/gofiber/fiber/v2"
	"os"
)

func cekPass() {
	//hash1, _ := bcrypt.GenerateFromPassword([]byte("Udin Gamboet"), bcrypt.DefaultCost)
	//fmt.Println(hash1)
	//fmt.Println(reflect.TypeOf(hash1))
	//
	//err := bcrypt.CompareHashAndPassword(hash1, []byte("Udin Gamboet"))
	//if err != nil {
	//	log.Println(err)
	//} else {
	//	log.Println("success")
	//}
}

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
