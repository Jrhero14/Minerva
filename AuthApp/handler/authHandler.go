package handler

import (
	"fmt"
	schemas "github.com/Jrhero14/Minerva/AuthApp/schemas"
	"github.com/Jrhero14/Minerva/config"
	"github.com/Jrhero14/Minerva/database"
	"github.com/Jrhero14/Minerva/database/model"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func CekRole(request *fiber.Ctx) bool {
	user := request.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	if claims["admin"].(bool) {
		return true
	} else {
		return false
	}
}

func Refresh(request *fiber.Ctx) error {
	cookie := request.Cookies("token")
	var err error
	token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Config("SECRET")), nil
	})
	if err != nil {
		return request.Status(400).JSON(fiber.Map{"status": "error", "Message": "Invalid Jwt or 'token' not found in cookies"})
	}
	payload := token.Claims.(jwt.MapClaims)
	claimsNew := jwt.MapClaims{
		"id":    payload["id"].(string),
		"name":  payload["name"].(string),
		"email": payload["email"].(string),
		"admin": payload["admin"].(bool),
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}
	// Create token
	tokenNew := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsNew)
	// Generate encoded token and send it as response.
	t, err := tokenNew.SignedString([]byte(config.Config("SECRET")))
	if err != nil {
		return request.SendStatus(fiber.StatusInternalServerError)
	}
	request.Set("token", t)
	request.Cookie(&fiber.Cookie{
		Name:  "token",
		Value: t,
	})
	return request.Status(200).JSON(fiber.Map{"Message": "refresh success"})
}

func Login(request *fiber.Ctx) error {
	db := database.DB
	userGet := new(model.User)
	loginBody := new(schemas.AuthSchema)
	err := request.BodyParser(&loginBody)
	if err != nil {
		return request.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	username := loginBody.Username
	password := loginBody.Password
	db.Preload("IdMember").Find(&userGet, "username = ?", username)
	if userGet.ID == uuid.Nil {
		return request.Status(400).JSON(fiber.Map{"status": "bad request", "message": "Wrong password or username"})
	}
	err = bcrypt.CompareHashAndPassword(userGet.Hash, []byte(password))
	if err != nil {
		return request.Status(400).JSON(fiber.Map{"status": "bad request", "message": "Wrong password or username", "data": err})
	} else {
		var admin bool
		if userGet.Role == "User" {
			admin = false
		} else {
			admin = true
		}
		claims := jwt.MapClaims{
			"id":    userGet.ID,
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

		request.Set("token", t)
		request.Cookie(&fiber.Cookie{
			Name:  "token",
			Value: t,
		})
		return request.Status(200).JSON(fiber.Map{"Message": "Success Login"})
	}
}

func AllUser(request *fiber.Ctx) error {
	if !CekRole(request) {
		return request.Status(403).JSON(fiber.Map{"status": "Forbidden", "message": "Hanya Admin atau Manager yang bisa melihat semua user"})
	}
	db := database.DB
	var users []model.User
	db.Preload("IdMember").Find(&users)
	if len(users) == 0 {
		return request.Status(404).JSON(fiber.Map{"status": "error", "message": "Tidak ada user", "data": users})
	}
	return request.JSON(fiber.Map{"status": "success", "message": "User ditemukan", "data": users})
}

func GetinfoMember(request *fiber.Ctx) error {
	db := database.DB
	var memberInfo model.Member
	var UserData model.User
	user := request.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	db.Preload("IdMember").Find(&UserData, "id = ?", claims["id"].(string))
	if UserData.ID == uuid.Nil {
		return request.Status(404).JSON(fiber.Map{"status": "error", "message": "User not Found"})
	}
	db.Find(&memberInfo, "id = ?", UserData.IdMember.ID)
	memberInfoSchema := struct {
		Nama      string
		BirthDay  string
		Institusi string
		Gender    string
		Alamat    string
		KodePos   string
		Email     string
		Phone     string
	}{
		Nama:      memberInfo.Nama,
		BirthDay:  memberInfo.BirthDay.Format("2006-01-02"),
		Institusi: memberInfo.Institusi,
		Gender:    memberInfo.Gender,
		Alamat:    memberInfo.Alamat,
		KodePos:   memberInfo.KodePos,
		Email:     memberInfo.Email,
		Phone:     memberInfo.Phone,
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
	bodyUser := new(schemas.RegSchema)

	err := request.BodyParser(&bodyUser)
	if err != nil {
		return request.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	userUUID := uuid.New()
	memberUUID := uuid.New()

	var favoriteBody model.Favorite
	favoriteBody.Id_Member = memberUUID
	err = db.Create(&favoriteBody).Error
	if err != nil {
		return request.Status(500).JSON(fiber.Map{"status": "error", "message": "Favorite Create Failed", "data": err})
	}

	var historyBody model.HistoryBorrow
	historyBody.IDMember = memberUUID
	err = db.Create(&historyBody).Error
	if err != nil {
		return request.Status(500).JSON(fiber.Map{"status": "error", "message": "History Create Failed", "data": err})
	}

	user.ID = userUUID
	member.ID = memberUUID
	member.Regis_date = time.Now()
	member.Exp_member = time.Now().AddDate(2, 0, 0)
	member.Id_Favorite = favoriteBody.ID
	member.IDFavorite = favoriteBody
	member.Id_HistoryBorrow = historyBody.ID
	member.IDHistory = historyBody
	err = db.Create(&member).Error
	if err != nil {
		return request.Status(500).JSON(fiber.Map{"status": "error", "message": "Member Create Failed", "data": err})
	}

	user.IdMem = member.ID
	user.Username = bodyUser.Username
	fmt.Println(bodyUser.Role)
	if i := bodyUser.Role; i == "1" {
		user.Role = "User"
	} else if i == "2" {
		user.Role = "Admin"
	} else {
		user.Role = "Manager"
	}
	hash, _ := bcrypt.GenerateFromPassword([]byte(bodyUser.Password), bcrypt.DefaultCost)
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
	var userMember model.User

	memberBody := new(schemas.MemberValueRegis)
	err := request.BodyParser(&memberBody)
	if err != nil {
		return request.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	user := request.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	idUser := claims["id"].(string)
	db.Preload("IdMember").Find(&userMember, "id = ?", idUser)
	db.Find(&member, "id = ?", userMember.IdMem)

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

	db.Save(&member)
	// Return the updated note
	return request.JSON(fiber.Map{"status": "success", "message": "Member Data Updated", "data": member})
}
