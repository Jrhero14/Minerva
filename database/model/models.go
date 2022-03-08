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
	Role             string
	Id_Favorite      int64
	IDFavorite       Favorite `gorm:"foreignkey:Id_Favorite;constraint:onUpdate:CASCADE,ondelete:SET NULL" json:"Favorite"`
	Id_HistoryBorrow int64
	IDHistory        HistoryBorrow `gorm:"foreignkey:Id_HistoryBorrow;constraint:onUpdate:CASCADE,ondelete:SET NULL" json:"History"`
	Id_Mylist        int64
}

type Mylist struct {
	gorm.Model
	ID       int64 `gorm:"primary_key:auto_increment"`
	NameList string
	IDMember uuid.UUID `gorm:"type:uuid"`
	Books    []*Book   `gorm:"many2many:mylist_books;" json:"Books"`
}

type HistoryBorrow struct {
	gorm.Model
	ID          int64         `gorm:"primary_key:auto_increment"`
	IDMember    uuid.UUID     `gorm:"type:uuid"`
	Prebookings []*PreBooking `gorm:"many2many:historyborrow_prebookings;" json:"Booked"`
}

type Favorite struct {
	gorm.Model
	ID        int64     `gorm:"primary_key:auto_increment"`
	Id_Member uuid.UUID `gorm:"type:uuid"`
	Books     []*Book   `gorm:"many2many:favorite_books;" json:"Books"`
}

// Book Entity
type Jenis struct {
	gorm.Model
	ID        int `gorm:"primary_key:auto_increment"`
	namaJenis string
}

type RakBuku struct {
	gorm.Model
	ID       int32 `gorm:"primary_key:auto_increment"`
	namaRak  string
	nomorRak string
	note     string `gorm:"type:text"`
}

type InfoDetail struct {
	gorm.Model
	ID        int32 `gorm:"primary_key:auto_increment"`
	Id_Book   int64
	NomorBuku string
	Id_Rak    int
	IDRak     RakBuku `gorm:"foreignkey:Id_Rak;constraint:onUpdate:CASCADE,ondelete:SET NULL" json:"RakBuku"`
}

type Kategori struct {
	gorm.Model
	ID           int32 `gorm:"primary_key:auto_increment"`
	NamaKategori string
}

type Book struct {
	gorm.Model
	ID            int64 `gorm:"primary_key:auto_increment"`
	Image         string
	Title         string
	JudulSeri     string
	Penerbit      string
	Deskripsi     string `gorm:"type:date"`
	Id_Jenis      int
	IDJenis       Jenis `gorm:"foreignkey:Id_Jenis;constraint:onUpdate:CASCADE,ondelete:CASCADE" json:"Jenis"`
	Bahasa        string
	ISBN          string
	Edisi         string
	Ketersediaan  bool
	Stock         int32
	Subjek        string
	Id_InfoDetail int32
	IDInfoDetail  InfoDetail `gorm:"foreignkey:Id_InfoDetail;constraint:onUpdate:CASCADE,ondelete:SET NULL" json:"InfoDetail"`
	Id_Kategori   int32
	IDKategori    Kategori `gorm:"foreignkey:Id_Kategori;constraint:onUpdate:CASCADE,ondelete:SET NULL" json:"Kategori"`
}

// Borrow Entity
type PreBooking struct {
	gorm.Model
	ID               int64 `gorm:"primary_key:auto_increment"`
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
	ID               int64 `gorm:"primary_key:auto_increment"`
	Id_DetailBook    int64
	IDInfoDetailBook InfoDetail `gorm:"foreignkey:Id_DetailBook;constraint:onUpdate:CASCADE,ondelete:SET NULL" json:"InfoDetail"`
	Id_Member        uuid.UUID  `gorm:"type:uuid"`
	IDMember         Member     `gorm:"foreignkey:Id_Member;constraint:onUpdate:CASCADE,ondelete:SET NULL" json:"Member"`
	Mobile           bool
	Borrowed         bool
	ExpireReturn     time.Time `gorm:"type:date"`
}
