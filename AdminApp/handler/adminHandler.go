package handler

import (
	"Minerva/database"
	"Minerva/database/model"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func CreateCategory(request *fiber.Ctx) error {
	db := database.DB
	newCategory := new(model.Kategori)
	err := request.BodyParser(&newCategory)
	if err != nil {
		return request.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	err = db.Create(&newCategory).Error
	if err != nil {
		return request.Status(500).JSON(fiber.Map{"status": "error", "message": "Can't Create Category", "data": err})
	}
	return request.Status(200).JSON(fiber.Map{"status": "success", "message": "Berhasil membuat kategori", "data": newCategory})
}

func GetAllCategory(request *fiber.Ctx) error {
	// Verify Admin
	user := request.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	if !claims["admin"].(bool) {
		return request.Status(403).JSON(fiber.Map{"message": "you not admin"})
	}

	db := database.DB
	var categories []model.Kategori

	db.Find(&categories)
	if len(categories) == 0 {
		return request.Status(404).JSON(fiber.Map{"status": "error", "message": "Tidak ada kategori", "data": categories})
		fmt.Println("WAHH GA ADA COEG")
	}
	return request.Status(200).JSON(fiber.Map{"status": "success", "message": "Kategori ditemukan", "data": categories})
}
