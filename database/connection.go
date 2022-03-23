package database

import (
	"Minerva/config"
	"Minerva/database/model"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"strconv"
	"time"
)

// Declare the variable for the database
var DB *gorm.DB

func CreateUserDefault() {
	db := DB
	uuIDMember1 := uuid.New()
	uuIDUser1 := uuid.New()
	var Member1 = model.Member{
		ID:         uuIDMember1,
		Nama:       "Admin Ganteng",
		Regis_date: time.Now(),
		Exp_member: time.Now().AddDate(5, 0, 0),
		Institusi:  "Minerva Librarian",
		Gender:     "L",
		Alamat:     "Denpasar",
		KodePos:    "12345678",
		Email:      "admin@gmail.com",
		Phone:      "08123456789",
	}
	db.Create(&Member1)
	hash1, _ := bcrypt.GenerateFromPassword([]byte("mimin123"), bcrypt.DefaultCost)
	var user1 = model.User{
		ID:       uuIDUser1,
		IdMem:    uuIDMember1,
		IdMember: Member1,
		Username: "adminganteng",
		Hash:     hash1,
		Role:     "2",
	}
	db.Create(&user1)

	// User Mobile
	uuIDMember2 := uuid.New()
	uuIDUser2 := uuid.New()
	var Favorite2 = model.Favorite{
		ID:        1,
		Id_Member: uuIDMember2,
	}
	db.Create(&Favorite2)
	var HistoryBorrow2 = model.HistoryBorrow{
		ID:       1,
		IDMember: uuIDMember2,
	}
	db.Create(&HistoryBorrow2)
	t1, err := time.Parse("2006-01-02", "1998-10-03")
	if err != nil {
		fmt.Println("Wahh date error")
	}
	var Member2 = model.Member{
		ID:               uuIDMember2,
		Nama:             "Udin Siawaludin",
		BirthDay:         t1,
		Regis_date:       time.Now(),
		Exp_member:       time.Now().AddDate(2, 0, 0),
		Institusi:        "Udayana",
		Gender:           "L",
		Alamat:           "Denpasar",
		KodePos:          "123456",
		Email:            "Didin14@gmail.com",
		Phone:            "081364368342",
		Id_Favorite:      Favorite2.ID,
		IDFavorite:       Favorite2,
		Id_HistoryBorrow: HistoryBorrow2.ID,
		IDHistory:        HistoryBorrow2,
	}
	db.Create(&Member2)
	hash2, _ := bcrypt.GenerateFromPassword([]byte("didin123"), bcrypt.DefaultCost)
	var user2 = model.User{
		ID:       uuIDUser2,
		IdMem:    uuIDMember2,
		IdMember: Member2,
		Username: "Udin14",
		Hash:     hash2,
		Role:     "1",
	}
	db.Create(&user2)
	fmt.Println("Create User default success")
}

func CreateJenis() {
	db := DB
	var Jenis1 = model.Jenis{
		ID:        1,
		NamaJenis: "Novel",
	}
	var Jenis2 = model.Jenis{
		ID:        2,
		NamaJenis: "Science",
	}
	var Jenis3 = model.Jenis{
		ID:        3,
		NamaJenis: "Komik",
	}
	var Jenis4 = model.Jenis{
		ID:        4,
		NamaJenis: "Ensiklopedia",
	}
	db.Create(&Jenis1)
	db.Create(&Jenis2)
	db.Create(&Jenis3)
	db.Create(&Jenis4)
}

func CreateCategory() {
	db := DB
	var Category1 = model.Kategori{
		ID:           1,
		NamaKategori: "Non Fiksi",
	}
	var Category2 = model.Kategori{
		ID:           2,
		NamaKategori: "Fiksi",
	}
	db.Create(&Category1)
	db.Create(&Category2)
}

func CreateRak() {
	db := DB
	var Rak1 = model.RakBuku{
		ID:       1,
		NamaRak:  "Rak Buku Novel 1",
		NomorRak: "N001",
		Note:     "Lantai 2 bagian fiksi",
	}
	var Rak2 = model.RakBuku{
		ID:       2,
		NamaRak:  "Rak Buku Science 1",
		NomorRak: "S001",
		Note:     "Lantai 2 bagian Non-fiksi",
	}
	var Rak3 = model.RakBuku{
		ID:       3,
		NamaRak:  "Rak Buku Komik 1",
		NomorRak: "K001",
		Note:     "Lantai 3 bagian fiksi",
	}
	db.Create(&Rak1)
	db.Create(&Rak2)
	db.Create(&Rak3)
}

func CreateBooks() {
	db := DB
	var jenis1 model.Jenis
	db.Find(&jenis1, "id = ?", 2)
	var kategori1 model.Kategori
	db.Find(&kategori1, "id = ?", 1)
	var buku1 = model.Book{
		ID:           1,
		Image:        "https://images-na.ssl-images-amazon.com/images/I/51Tp9VTd6QL._SX331_BO1,204,203,200_.jpg",
		Title:        "Linux In A Nutshell",
		JudulSeri:    "In a Nutshell",
		Penerbit:     "Sebastopol, CA : OReilly., 2005",
		Deskripsi:    "Linux dalam kulit kacang goreng",
		Id_Jenis:     2,
		IDJenis:      jenis1,
		Bahasa:       "English",
		ISBN:         "9780596009304",
		Edisi:        "Edisi ke-6",
		Ketersediaan: false,
		Stock:        0,
		Subjek:       "Operation System",
		Id_Kategori:  1,
		IDKategori:   kategori1,
	}
	db.Create(&buku1)

	// Book 2
	var jenis2 model.Jenis
	db.Find(&jenis2, "id = ?", 1)
	var kategori2 model.Kategori
	db.Find(&kategori2, "id = ?", 2)
	var buku2 = model.Book{
		ID:           2,
		Image:        "https://opac.perpusnas.go.id/uploaded_files/sampul_koleksi/original/Monograf/1098530.jpg?rnd=1291746925",
		Title:        "Dunia Sophie",
		JudulSeri:    "Sebuah novel filsafat",
		Penerbit:     "Bandung : Penerbit Mizan, 2018",
		Deskripsi:    "Sophie, seorang pelajar sekolah menengah berusia empat belas tahun. Suatu hari pulang sekolah, dia mendapat sebuah surat misterius yang hanya berisikan satu pertanyaan...",
		Id_Jenis:     1,
		IDJenis:      jenis2,
		Bahasa:       "Indonesia",
		ISBN:         "9786024410209",
		Edisi:        "Edisi Ketiga, Cetakan IV",
		Ketersediaan: false,
		Stock:        0,
		Subjek:       "Filsafat",
		Id_Kategori:  2,
		IDKategori:   kategori1,
	}
	db.Create(&buku2)
}

func CreateStock() {
	db := DB
	var book1 model.Book
	var book2 model.Book
	db.Find(&book1, "id = ?", 1)
	db.Find(&book2, "id = ?", 2)
	var rakBuku1 model.RakBuku
	var rakBuku2 model.RakBuku
	db.Find(&rakBuku1, "id = ?", 2)
	db.Find(&rakBuku2, "id = ?", 1)
	var stock1 = model.InfoDetail{
		ID:        1,
		Id_Book:   1,
		NomorBuku: "LN1",
		Ready:     true,
		Id_Rak:    2,
		IDRak:     rakBuku1,
	}
	var stock2 = model.InfoDetail{
		ID:        2,
		Id_Book:   1,
		NomorBuku: "LN2",
		Ready:     true,
		Id_Rak:    2,
		IDRak:     rakBuku1,
	}
	var stock3 = model.InfoDetail{
		ID:        3,
		Id_Book:   1,
		NomorBuku: "LN3",
		Ready:     true,
		Id_Rak:    2,
		IDRak:     rakBuku1,
	}
	var stock4 = model.InfoDetail{
		ID:        4,
		Id_Book:   2,
		NomorBuku: "DS1",
		Ready:     true,
		Id_Rak:    1,
		IDRak:     rakBuku2,
	}
	var stock5 = model.InfoDetail{
		ID:        5,
		Id_Book:   2,
		NomorBuku: "DS2",
		Ready:     true,
		Id_Rak:    1,
		IDRak:     rakBuku2,
	}
	var stock6 = model.InfoDetail{
		ID:        6,
		Id_Book:   2,
		NomorBuku: "DS3",
		Ready:     true,
		Id_Rak:    1,
		IDRak:     rakBuku2,
	}
	db.Create(&stock1)
	db.Create(&stock2)
	db.Create(&stock3)
	db.Create(&stock4)
	db.Create(&stock5)
	db.Create(&stock6)

	// update stock
	var bookGet1 model.Book
	db.Find(&bookGet1, "id = ?", 1).Preload("IDJenis").Preload("IDKategori")
	bookGet1.Stock = 3
	bookGet1.Ketersediaan = true
	db.Save(&bookGet1)

	var bookGet2 model.Book
	db.Find(&bookGet2, "id = ?", 2).Preload("IDJenis").Preload("IDKategori")
	bookGet2.Stock = 3
	bookGet2.Ketersediaan = true
	db.Save(&bookGet2)
}

func CreateBoooked() {
	db := DB
	var bookStock1 model.InfoDetail
	db.Preload("IDRak").Find(&bookStock1, "id = ?", 3)
	bookStock1.Ready = false
	db.Save(&bookStock1)
	var bookStock2 model.InfoDetail
	db.Preload("IDRak").Find(&bookStock2, "id = ?", 4)
	bookStock2.Ready = false
	db.Save(&bookStock2)
	var member model.Member
	db.Find(&member, "nama = ?", "Udin Siawaludin")
	var historyborrow model.HistoryBorrow
	db.Find(&historyborrow, "id = ?", 1)
	var prebookeds = []model.PreBooking{
		model.PreBooking{
			ID:               1,
			IDBook:           bookStock1.Id_Book,
			Id_DetailBook:    int64(bookStock1.ID),
			IDInfoDetailBook: bookStock1,
			Id_Member:        member.ID,
			IDMember:         member,
			Mobile:           true,
			Borrowed:         false,
			ExpireBorrow:     time.Now().AddDate(0, 0, 7),
		},
		{
			ID:               2,
			IDBook:           bookStock2.Id_Book,
			Id_DetailBook:    int64(bookStock2.ID),
			IDInfoDetailBook: bookStock2,
			Id_Member:        member.ID,
			IDMember:         member,
			Mobile:           true,
			Borrowed:         false,
			ExpireBorrow:     time.Now().AddDate(0, 0, 7),
		},
	}
	db.Create(&prebookeds)
}

func CreateDummy() {
	CreateUserDefault()
	CreateJenis()
	CreateCategory()
	CreateRak()
	CreateBooks()
	CreateStock()
	CreateBoooked()
}

// ConnectDB connect to db
func ConnectDB() {
	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		log.Println("Idiot")
	}

	// Connection URL to connect to Postgres Database
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))
	// Connect to the DB and initialize the DB variable
	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")

	// Migrate the database
	//DB.AutoMigrate(&model.Book{})
	//DB.AutoMigrate(&model.User{})
	//DB.AutoMigrate(&model.Member{})
	//DB.AutoMigrate(&model.RakBuku{})
	//DB.AutoMigrate(&model.Kategori{})
	//DB.AutoMigrate(&model.InfoDetail{})
	//DB.AutoMigrate(&model.Jenis{})
	//DB.AutoMigrate(&model.PreBooking{})
	//DB.AutoMigrate(&model.Booked{})
	//fmt.Println("Database Migrated")
	//CreateDummy()
	//fmt.Println("Data Dummy Created")
}
