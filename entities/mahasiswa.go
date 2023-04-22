package entities

import (
	"time"
)

type Mahasiswa struct {
	Id                uint
	Nama              string
	Usia              uint
	Gender            uint
	TanggalRegistrasi time.Time
}

type MahasiswaHobi struct {
	IdMahasiswa uint
	IdHobi      uint
}
