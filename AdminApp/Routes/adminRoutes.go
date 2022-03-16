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
	// Create Jenis
	admin.Post("/create-jenis", adminHandler.CreateJenis)
	// Get All Jenis
	admin.Get("/all-jenis", adminHandler.GetAllJenis)
	// Create Rak Buku
	admin.Post("/create-rak", adminHandler.CreateRak)
	// Get All Rak buku
	admin.Get("/all-rak", adminHandler.GetAllRak)
	// Create new Book
	admin.Post("/create-book", adminHandler.CreateNewBook)
	// Restock Book
	admin.Post("/restock", adminHandler.RestockBook)
	// All Books
	admin.Get("all-book", adminHandler.AllBooks)
}
