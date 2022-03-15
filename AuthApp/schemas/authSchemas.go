package schemas

import "github.com/google/uuid"

type MemberValueRegis struct {
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
	IdMember MemberValueRegis
	Username string
}

type AuthSchema struct {
	Username string
	Password string
}

type RegSchema struct {
	Username string
	Password string
	Role     string
}
