package requests

type CreateMahasiswaRequest struct {
	Nama    string   `json:"nama" form:"nama" query:"nama" validate:"required"`
	Usia    uint      `json:"usia" form:"usia" query:"usia" validate:"required"`
	Gender  uint      `json:"gender" form:"gender" query:"gender" validate:"required"`
	Jurusan []string `json:"jurusan" form:"jurusan" query:"jurusan" validate:"required"`
	Hobi    []string `json:"hobi" form:"hobi" query:"hobi"`
}
