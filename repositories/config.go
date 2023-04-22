package repositories

import (
	"database/sql"
	"main/entities"
	"main/entities/responses"
)

type repositories struct {
	db *sql.DB
}

type Repositories interface {
	GetAllMahasiswa() ([]responses.MahasiswaResponse, error)
	GetMahasiswaById(id uint) (responses.MahasiswaResponse, error)
	GetMahasiswaByName(name string) ([]responses.MahasiswaResponse, error)
	InsertNewMahasiswa(mahasiswa entities.Mahasiswa) (entities.Mahasiswa, error)
	InserNewMahasiswaHobi(mahasiswaId uint, hobiIds []uint) error
	InserNewMahasiswaJurusan(mahasiswaId uint, jurusanIds []uint) error
	UpdateMahasiswa(mahasiswa entities.Mahasiswa) error
	DeleteMahasiswaById(id uint) error

	GetAllHobi() ([]entities.Hobi, error)
	InsertNewHobi(hobi []string) ([]uint, error)

	GetAllJurusan() ([]entities.Jurusan, error)
	InsertNewJurusan(jurusan []string) ([]uint, error)
}

func NewRepositories(db *sql.DB) Repositories {
	return &repositories{db: db}
}
