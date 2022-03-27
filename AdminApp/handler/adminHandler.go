package handler

import (
	schemas "github.com/Jrhero14/Minerva/AdminApp/schemas"
	auth "github.com/Jrhero14/Minerva/AuthApp/handler"
	schemas2 "github.com/Jrhero14/Minerva/AuthApp/schemas"
	"github.com/Jrhero14/Minerva/database"
	"github.com/Jrhero14/Minerva/database/model"
	"github.com/gofiber/fiber/v2"
)

func CreateCategory(request *fiber.Ctx) error {
	if !auth.CekRole(request) {
		return request.Status(403).JSON(fiber.Map{"status": "Forbidden", "message": "Hanya Admin atau Manager yang bisa mengakses"})
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
		return request.Status(403).JSON(fiber.Map{"status": "Forbidden", "message": "Hanya Admin atau Manager yang bisa mengakses"})
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
		return request.Status(403).JSON(fiber.Map{"status": "Forbidden", "message": "Hanya Admin atau Manager yang bisa mengakses"})
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
		return request.Status(403).JSON(fiber.Map{"status": "Forbidden", "message": "Hanya Admin atau Manager yang bisa mengakses"})
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
		return request.Status(403).JSON(fiber.Map{"status": "Forbidden", "message": "Hanya Admin atau Manager yang bisa mengakses"})
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
		return request.Status(403).JSON(fiber.Map{"status": "Forbidden", "message": "Hanya Admin atau Manager yang bisa mengakses"})
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
	if !auth.CekRole(request) {
		return request.Status(403).JSON(fiber.Map{"status": "Forbidden", "message": "Hanya Admin atau Manager yang bisa mengakses"})
	}
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

func AllBooks(request *fiber.Ctx) error {
	if !auth.CekRole(request) {
		return request.Status(403).JSON(fiber.Map{"status": "Forbidden", "message": "Hanya Admin atau Manager yang bisa mengakses"})
	}
	db := database.DB
	var Books []model.Book
	err := db.Find(&Books).Error
	if err != nil {
		return request.Status(404).JSON(fiber.Map{"status": "error", "message": "not found books"})
	}
	return request.Status(200).JSON(fiber.Map{"status": "sucess", "message": "found books", "data": Books})
}

func DetailBook(request *fiber.Ctx) error {
	db := database.DB
	var bodyQuery schemas.DetailBook
	var Book model.Book
	err := request.BodyParser(&bodyQuery)
	if err != nil {
		return request.Status(400).JSON(fiber.Map{"Message": "Review your input"})
	}
	db.Preload("IDJenis").Preload("IDKategori").Find(&Book, "id = ?", bodyQuery.IdBook)
	if Book.Id_Jenis == 0 {
		return request.Status(400).JSON(fiber.Map{"Message": "Can't Found book"})
	}
	return request.Status(200).JSON(fiber.Map{"Message": "Success", "data": Book})
}

func RestockBook(request *fiber.Ctx) error {
	if !auth.CekRole(request) {
		return request.Status(403).JSON(fiber.Map{"status": "Forbidden", "message": "Hanya Admin atau Manager yang bisa mengakses"})
	}
	db := database.DB
	bodyRestock := new(schemas.Restock)
	restock := new(model.InfoDetail)
	var Book model.Book

	err := request.BodyParser(&bodyRestock)
	if err != nil {
		return request.Status(400).JSON(fiber.Map{"status": "error", "message": "Review your input"})
	}
	var rak model.RakBuku
	err = db.Find(&rak, "id = ?", bodyRestock.IdRak).Error
	if err != nil {
		return request.Status(404).JSON(fiber.Map{"status": "error", "message": "Can't find Rak Buku"})
	}
	err = db.Preload("IDJenis").Preload("IDKategori").Find(&Book, "id = ?", bodyRestock.IdBook).Error
	if err != nil {
		return request.Status(404).JSON(fiber.Map{"status": "error", "message": "Can't find Buku"})
	}
	restock.Id_Book = bodyRestock.IdBook
	restock.NomorBuku = bodyRestock.NomorBuku
	restock.Ready = true
	restock.Id_Rak = bodyRestock.IdRak
	restock.IDRak = rak
	err = db.Create(&restock).Error
	if err != nil {
		return request.Status(500).JSON(fiber.Map{"status": "error", "message": "Can't create new Stock"})
	}
	Book.Stock += 1
	Book.Ketersediaan = true
	err = db.Save(&Book).Error
	if err != nil {
		return request.Status(500).JSON(fiber.Map{"status": "error", "message": "Can't Save Book"})
	}

	return request.Status(200).JSON(fiber.Map{"status": "success", "message": "yey new stock", "data": restock})

}

func GetBookStock(request *fiber.Ctx) error {
	db := database.DB
	var stocks []model.InfoDetail
	var body struct {
		IdBook int
	}

	err := request.BodyParser(&body)
	if err != nil {
		return request.Status(400).JSON(fiber.Map{"status": "error", "message": "Review your input"})
	}
	db.Preload("IDRak").Find(&stocks, "Id_Book = ?", body.IdBook)
	if len(stocks) == 0 {
		return request.Status(500).JSON(fiber.Map{"status": "error", "message": "stock buku tersebut kosong"})
	}
	return request.Status(200).JSON(fiber.Map{"status": "success", "totalData": len(stocks), "data": stocks})
}

func GetHistoryBooked(request *fiber.Ctx) error {
	db := database.DB
	var history []model.PreBooking
	var body schemas2.BodyBorrowsBook
	err := request.BodyParser(&body)
	if err != nil {
		return request.Status(400).JSON(fiber.Map{"status": "error", "message": "review your input body"})
	}
	db.Preload("IDMember").Find(&history, "id_book = ?", body.IdBook)
	return request.Status(200).JSON(fiber.Map{"data": history, "total": len(history)})
}

//func Booking(request *fiber.Ctx) error {
//	db := database.DB
//	Booking := new(model.PreBooking)
//	var BookAvailable []model.InfoDetail
//
//}
