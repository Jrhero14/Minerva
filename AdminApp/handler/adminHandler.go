package handler

import (
	schemas "Minerva/AdminApp/schemas"
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

func CreateNewBook(request *fiber.Ctx) error {
	db := database.DB
	BookNew := new(model.Book)
	bodyBook := new(schemas.BookBody)
	var jenis model.Jenis
	var kategory model.Kategori
	err := request.BodyParser(&bodyBook)
	if err != nil {
		return request.Status(400).JSON(fiber.Map{"status": "error", "messsage": "review your input"})
	}
	err = db.Find(&jenis, "id = ?", bodyBook.IdJenis).Error
	if err != nil {
		return request.Status(500).JSON(fiber.Map{"status": "error", "messsage": "Can't find jenis"})
	}
	err = db.Find(&kategory, "id = ?", bodyBook.IdKategori).Error
	if err != nil {
		return request.Status(500).JSON(fiber.Map{"status": "error", "messsage": "Can't find kategori"})
	}
	BookNew.Image = bodyBook.Image
	BookNew.Title = bodyBook.Title
	BookNew.JudulSeri = bodyBook.JudulSeri
	BookNew.Penerbit = bodyBook.Penerbit
	BookNew.Deskripsi = bodyBook.Deskripsi
	BookNew.Id_Jenis = bodyBook.IdJenis
	BookNew.IDJenis = jenis
	BookNew.Bahasa = bodyBook.Bahasa
	BookNew.ISBN = bodyBook.ISBN
	BookNew.Edisi = bodyBook.Edisi
	BookNew.Subjek = bodyBook.Subjek
	BookNew.Id_Kategori = bodyBook.IdKategori
	BookNew.IDKategori = kategory
	BookNew.Ketersediaan = false
	BookNew.Stock = 0

	err = db.Create(&BookNew).Error
	if err != nil {
		return request.Status(500).JSON(fiber.Map{"status": "error", "messsage": "Can't create new book"})
	}
	return request.Status(200).JSON(fiber.Map{"status": "success", "messsage": "Success Create Book", "data": BookNew})
}
