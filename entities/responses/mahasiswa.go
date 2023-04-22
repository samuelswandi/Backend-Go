package responses

import (
	"time"
)

type MahasiswaResponse struct {
	Id                uint
	Nama              string
	Usia              int
	Gender            int
	TanggalRegistrasi time.Time
	Jurusan           string
	Hobi              string
}
