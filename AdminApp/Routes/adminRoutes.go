package Routes

import (
	adminHandler "github.com/Jrhero14/Minerva/AdminApp/handler"
	"github.com/Jrhero14/Minerva/config"
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
	admin.Get("/all-book", adminHandler.AllBooks)
	// Get Book Stocks
	admin.Post("/book-stocks", adminHandler.GetBookStock)
	// get History Borrow Book
	admin.Post("/history", adminHandler.GetHistoryBooked)
	// Get Detail One Book
	admin.Post("/detail-book", adminHandler.DetailBook)
}
