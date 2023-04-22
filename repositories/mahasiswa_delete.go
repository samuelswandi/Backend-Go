package repositories

func (repo *repositories) DeleteMahasiswaById(id uint) error {
	db := repo.db

	tx, _ := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// hapus dari mahasiswa_hobi
	_, err := db.Query("DELETE from mahasiswa_hobi WHERE id_mahasiswa = ?", id)
	if err != nil {
		return err
	}

	// hapus dari mahasiswa_jurusan
	_, err = db.Query("DELETE from mahasiswa_jurusan WHERE id_mahasiswa = ?", id)
	if err != nil {
		return err
	}

	// hapus dari mahasiswa
	_, err = db.Query("DELETE from mahasiswa WHERE id = ?", id)
	if err != nil {
		return err
	}

	tx.Commit()
	return nil
}
