package requests

type UpdateMahasiswaRequest struct {
	Id     uint   `json:"id" form:"id" query:"id" validate:"required"`
	Nama   string `json:"nama" form:"nama" query:"nama" validate:"required"`
	Usia   uint   `json:"usia" form:"usia" query:"usia" validate:"required"`
	Gender uint   `json:"gender" form:"gender" query:"gender" validate:"required"`
}
