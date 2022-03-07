package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid"`
	IdMem    uuid.UUID `gorm:"not null"`
	IdMember Member    `gorm:"foreignkey:IdMem;constraint:onUpdate:CASCADE,ondelete:CASCADE" json:"member"`
	Username string
	Hash     []uint8
}

type Member struct {
	ID         uuid.UUID
	Nama       string
	BirthDay   time.Time `gorm:"type:date"`
	Regis_date time.Time `gorm:"type:date"`
	Exp_member time.Time `gorm:"type:date"`
	Institusi  string
	Gender     string
	Alamat     string `gorm:"type:text"`
	KodePos    string
	Email      string
	Phone      string
	Role       string
	//id_favorite
	//id_history
	//id_mylist
}
