package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// User Entity
type User struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid"`
	IdMem    uuid.UUID `gorm:"not null"`
	IdMember Member    `gorm:"foreignkey:IdMem;constraint:onUpdate:CASCADE,ondelete:CASCADE" json:"member"`
	Username string
	Hash     []uint8
	Role     string
}

type Member struct {
	gorm.Model
	ID               uuid.UUID
	Nama             string
	BirthDay         time.Time `gorm:"type:date"`
	Regis_date       time.Time `gorm:"type:date"`
	Exp_member       time.Time `gorm:"type:date"`
	Institusi        string
	Gender           string
	Alamat           string `gorm:"type:text"`
	KodePos          string
	Email            string
	Phone            string
	Id_Favorite      int64         `gorm:"default:null"`
	IDFavorite       Favorite      `gorm:"foreignkey:Id_Favorite;constraint:onUpdate:CASCADE,ondelete:SET NULL" json:"Favorite"`
	Id_HistoryBorrow int64         `gorm:"default:null"`
	IDHistory        HistoryBorrow `gorm:"foreignkey:Id_HistoryBorrow;constraint:onUpdate:CASCADE,ondelete:SET NULL" json:"History"`
	Id_Mylist        int64         `gorm:"default:null"`
}

type Mylist struct {
	gorm.Model
	ID       int64 `gorm:"primary_key:autoIncrement"`
	NameList string
	IDMember uuid.UUID `gorm:"type:uuid"`
	Books    []*Book   `gorm:"many2many:mylist_books;" json:"Books"`
}

type HistoryBorrow struct {
	gorm.Model
	ID          int64         `gorm:"primary_key:autoIncrement"`
	IDMember    uuid.UUID     `gorm:"type:uuid"`
	Prebookings []*PreBooking `gorm:"many2many:historyborrow_prebookings;" json:"Booked"`
}

type Favorite struct {
	gorm.Model
	ID        int64     `gorm:"primary_key:autoIncrement"`
	Id_Member uuid.UUID `gorm:"type:uuid"`
	Books     []*Book   `gorm:"many2many:favorite_books;" json:"Books"`
}

// Book Entity
type Jenis struct {
	gorm.Model
	ID        int `gorm:"primary_key:autoIncrement"`
	NamaJenis string
}

type RakBuku struct {
	gorm.Model
	ID       int32 `gorm:"primary_key:autoIncrement"`
	NamaRak  string
	NomorRak string
	Note     string `gorm:"type:text"`
}

type InfoDetail struct {
	gorm.Model
	ID        int32 `gorm:"primaryKey;autoIncrement" json:"id"`
	Id_Book   int64
	NomorBuku string
	Id_Rak    int
	Ready     bool
	IDRak     RakBuku `gorm:"foreignkey:Id_Rak;constraint:onUpdate:CASCADE,ondelete:SET NULL" json:"RakBuku"`
}

type Kategori struct {
	gorm.Model
	ID           int32 `gorm:"primary_key:autoIncrement"`
	NamaKategori string
}

type Book struct {
	gorm.Model
	ID           int64 `gorm:"primary_key:autoIncrement"`
	Image        string
	Title        string
	JudulSeri    string
	Penerbit     string
	Deskripsi    string `gorm:"type:text"`
	Id_Jenis     int
	IDJenis      Jenis `gorm:"foreignkey:Id_Jenis;constraint:onUpdate:CASCADE,ondelete:CASCADE" json:"Jenis"`
	Bahasa       string
	ISBN         string
	Edisi        string
	Ketersediaan bool
	Stock        int32
	Subjek       string
	Id_Kategori  int32
	IDKategori   Kategori `gorm:"foreignkey:Id_Kategori;constraint:onUpdate:CASCADE,ondelete:SET NULL" json:"Kategori"`
}

// Borrow Entity
type PreBooking struct {
	gorm.Model
	ID               int64 `gorm:"primary_key:autoIncrement"`
	IDBook           int64
	Id_DetailBook    int64
	IDInfoDetailBook InfoDetail `gorm:"foreignkey:Id_DetailBook;constraint:onUpdate:CASCADE,ondelete:SET NULL" json:"InfoDetail"`
	Id_Member        uuid.UUID  `gorm:"type:uuid"`
	IDMember         Member     `gorm:"foreignkey:Id_Member;constraint:onUpdate:CASCADE,ondelete:SET NULL" json:"Member"`
	Mobile           bool
	Borrowed         bool
	ExpireBorrow     time.Time `gorm:"type:date"`
}

type Booked struct {
	gorm.Model
	ID               int64 `gorm:"primary_key:autoIncrement"`
	Id_DetailBook    int64
	IDInfoDetailBook InfoDetail `gorm:"foreignkey:Id_DetailBook;constraint:onUpdate:CASCADE,ondelete:SET NULL" json:"InfoDetail"`
	Id_Member        uuid.UUID  `gorm:"type:uuid"`
	IDMember         Member     `gorm:"foreignkey:Id_Member;constraint:onUpdate:CASCADE,ondelete:SET NULL" json:"Member"`
	Mobile           bool
	Borrowed         bool
	ExpireReturn     time.Time `gorm:"type:date"`
}
