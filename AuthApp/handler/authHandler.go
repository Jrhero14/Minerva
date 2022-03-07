package handler

import (
	"Minerva/config"
	"Minerva/database"
	"Minerva/database/model"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

type memberValueRegis struct {
	ID        string
	Nama      string
	BirthDay  string
	Institusi string
	Gender    string
	Alamat    string
	KodePos   string
	Email     string
	Phone     string
	Role      string
}

type User struct {
	ID       uuid.UUID
	IdMem    uuid.UUID
	IdMember memberValueRegis
	Username string
}

func Login(request *fiber.Ctx) error {
	db := database.DB
	userGet := new(model.User)
	username := request.FormValue("Username")
	password := request.FormValue("Password")
	db.Preload("IdMember").Find(&userGet, "username = ?", username)
	if userGet.ID == uuid.Nil {
		return request.Status(404).JSON(fiber.Map{"status": "error", "message": "user not found", "data": userGet})
	}
	err := bcrypt.CompareHashAndPassword(userGet.Hash, []byte(password))
	if err != nil {
		log.Println(err)
	} else {
		var admin bool
		if userGet.IdMember.Role == "regular" {
			admin = false
		} else {
			admin = true
		}
		claims := jwt.MapClaims{
			"name":  userGet.Username,
			"email": userGet.IdMember.Email,
			"admin": admin,
			"exp":   time.Now().Add(time.Hour * 24).Unix(),
		}
		// Create token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte(config.Config("SECRET")))
		if err != nil {
			return request.SendStatus(fiber.StatusInternalServerError)
		}

		return request.JSON(fiber.Map{"token": t})
	}
	return request.JSON(fiber.Map{"status": "success", "message": "User ditemukan", "data": userGet})
}

func AllUser(request *fiber.Ctx) error {
	user := request.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	fmt.Println(name)
	db := database.DB
	var users []model.User
	db.Preload("IdMember").Find(&users)
	if len(users) == 0 {
		request.Status(404).JSON(fiber.Map{"status": "error", "message": "Tidak ada user", "data": users})
	}
	return request.JSON(fiber.Map{"status": "success", "message": "User ditemukan", "data": users})
}

func GetinfoMember(request *fiber.Ctx) error {
	db := database.DB
	var memberInfo model.Member
	db.Find(&memberInfo, "id = ?", request.FormValue("idMem"))
	if memberInfo.ID == uuid.Nil {
		request.Status(404).JSON(fiber.Map{"status": "error", "message": "Member Info not Found", "data": memberInfo})
	}
	memberInfoSchema := struct {
		Nama      string
		BirthDay  string
		Institusi string
		Gender    string
		Alamat    string
		KodePos   string
		Email     string
		Phone     string
		Role      string
	}{
		Nama:      memberInfo.Nama,
		BirthDay:  memberInfo.BirthDay.Format("2006-01-02"),
		Institusi: memberInfo.Institusi,
		Gender:    memberInfo.Gender,
		Alamat:    memberInfo.Alamat,
		KodePos:   memberInfo.KodePos,
		Email:     memberInfo.Email,
		Phone:     memberInfo.Phone,
		Role:      memberInfo.Role,
	}
	return request.JSON(fiber.Map{
		"status":  "success",
		"message": "member info found",
		"data":    memberInfoSchema,
	})
}

func CreateUser(request *fiber.Ctx) error {
	db := database.DB
	user := new(model.User)
	member := new(model.Member)

	user.ID = uuid.New()
	member.ID = uuid.New()
	member.Regis_date = time.Now()
	member.Exp_member = time.Now().AddDate(2, 0, 0)
	err := db.Create(&member).Error
	if err != nil {
		return request.Status(500).JSON(fiber.Map{"status": "error", "message": "Member Create Failed", "data": err})
	}

	user.IdMem = member.ID
	user.Username = request.FormValue("username")
	hash, _ := bcrypt.GenerateFromPassword([]byte(request.FormValue("password")), bcrypt.DefaultCost)
	user.Hash = hash
	err = db.Create(&user).Error
	if err != nil {
		return request.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create user", "data": err})
	}

	// Return the created note
	return request.JSON(fiber.Map{"status": "success", "message": "Created User", "data": user})
}

func UpdateMember(request *fiber.Ctx) error {
	db := database.DB
	var member model.Member

	memberBody := new(memberValueRegis)
	err := request.BodyParser(&memberBody)
	if err != nil {
		return request.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	db.Find(&member, "id = ?", memberBody.ID)

	t1, err := time.Parse("2006-01-02", memberBody.BirthDay)
	if err != nil {
		fmt.Println("Wahh date error")
	}

	member.Nama = memberBody.Nama
	member.BirthDay = t1
	member.Institusi = memberBody.Institusi
	member.Gender = memberBody.Gender
	member.Alamat = memberBody.Alamat
	member.KodePos = memberBody.KodePos
	member.Email = memberBody.Email
	member.Phone = memberBody.Phone
	member.Role = memberBody.Role

	db.Save(&member)
	// Return the updated note
	return request.JSON(fiber.Map{"status": "success", "message": "Member Data Updated", "data": member})
}
