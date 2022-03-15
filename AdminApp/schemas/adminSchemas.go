package schemas

type BookBody struct {
	Image        string
	Title        string
	JudulSeri    string
	Penerbit     string
	Deskripsi    string
	IdJenis      int
	Bahasa       string
	ISBN         string
	Edisi        string
	Ketersediaan bool
	Stock        int32
	Subjek       string
	NomorBuku    string
	IdRak        int
	IdKategori   int32
}
