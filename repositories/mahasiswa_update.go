package repositories

import "main/entities"

func (repo *repositories) UpdateMahasiswa(mahasiswa entities.Mahasiswa) error {
	db := repo.db

	_, err := db.Query("UPDATE mahasiswa SET nama = ?, usia = ?, gender = ? WHERE id = ?", mahasiswa.Nama, mahasiswa.Usia, mahasiswa.Gender, mahasiswa.Id)
	if err != nil {
		return err
	}

	return nil
}
