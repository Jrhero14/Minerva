package handler

import (
	auth "Minerva/AuthApp/handler"
	"Minerva/database"
	"Minerva/database/model"
	"github.com/gofiber/fiber/v2"
)

func CreateCategory(request *fiber.Ctx) error {
	if !auth.CekRole(request) {
		return request.Status(403).JSON(fiber.Map{"status": "Forbidden", "message": "Hanya Admin atau Manager yang bisa melihat semua user"})
	}
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
	if !auth.CekRole(request) {
		return request.Status(403).JSON(fiber.Map{"status": "Forbidden", "message": "Hanya Admin atau Manager yang bisa melihat semua user"})
	}

	db := database.DB
	var categories []model.Kategori

	db.Find(&categories)
	if len(categories) == 0 {
		return request.Status(404).JSON(fiber.Map{"status": "error", "message": "Tidak ada kategori", "data": categories})
	}
	return request.Status(200).JSON(fiber.Map{"status": "success", "message": "Kategori ditemukan", "data": categories})
}

func CreateJenis(request *fiber.Ctx) error {
	if !auth.CekRole(request) {
		return request.Status(403).JSON(fiber.Map{"status": "Forbidden", "message": "Hanya Admin atau Manager yang bisa melihat semua user"})
	}
	db := database.DB
	newJenis := new(model.Jenis)
	err := request.BodyParser(&newJenis)
	if err != nil {
		return request.Status(400).JSON(fiber.Map{"status": "error", "message": "Data input salah, periksa kembali", "data": err})
	}
	err = db.Create(&newJenis).Error
	if err != nil {
		return request.Status(500).JSON(fiber.Map{"status": "error", "message": "Tidak bisa membuat Jenis baru", "data": err})
	}
	return request.Status(200).JSON(fiber.Map{"status": "success", "message": "Berhasil membuat jenis baru", "data": newJenis})
}

func GetAllJenis(request *fiber.Ctx) error {
	if !auth.CekRole(request) {
		return request.Status(403).JSON(fiber.Map{"status": "Forbidden", "message": "Hanya Admin atau Manager yang bisa melihat semua user"})
	}
	db := database.DB
	var allJenis []model.Jenis
	db.Find(&allJenis)
	if len(allJenis) == 0 {
		return request.Status(404).JSON(fiber.Map{"status": "not found", "data": allJenis})
	}
	return request.Status(200).JSON(fiber.Map{"status": "success", "message": "Data ditemukan", "data": allJenis})
}

func CreateRak(request *fiber.Ctx) error {
	if !auth.CekRole(request) {
		return request.Status(403).JSON(fiber.Map{"status": "Forbidden", "message": "Hanya Admin atau Manager yang bisa melihat semua user"})
	}
	db := database.DB
	rakbukuBody := new(model.RakBuku)
	err := request.BodyParser(&rakbukuBody)
	if err != nil {
		return request.Status(402).JSON(fiber.Map{"status": "error", "message": "Review your input data", "data": err})
	}
	err = db.Create(&rakbukuBody).Error
	if err != nil {
		return request.Status(500).JSON(fiber.Map{"status": "error", "message": "Can't Create new Rak Buku", "data": err})
	}
	return request.Status(200).JSON(fiber.Map{"status": "success", "message": "Create new Rak success", "data": rakbukuBody})
}

func GetAllRak(request *fiber.Ctx) error {
	if !auth.CekRole(request) {
		return request.Status(403).JSON(fiber.Map{"status": "Forbidden", "message": "Hanya Admin atau Manager yang bisa melihat semua user"})
	}
	db := database.DB
	var allRak []model.RakBuku
	db.Find(&allRak)
	if len(allRak) == 0 {
		return request.Status(404).JSON(fiber.Map{"status": "error", "message": "Not Found Rak Buku", "data": allRak})
	}
	return request.Status(200).JSON(fiber.Map{"status": "success", "message": "Rak Buku found", "data": allRak})
}
