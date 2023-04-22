package repositories

import (
	"main/entities"
)

func (repo *repositories) GetAllJurusan() ([]entities.Jurusan, error) {
	db := repo.db

	rows, err := db.Query("select * from jurusan")
	if err != nil {
		return []entities.Jurusan{}, err
	}
	defer rows.Close()

	var result []entities.Jurusan

	for rows.Next() {
		var each = entities.Jurusan{}
		var err = rows.Scan(&each.Id, &each.Nama)

		if err != nil {
			return []entities.Jurusan{}, err
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		return []entities.Jurusan{}, err
	}

	return result, nil
}

func (repo *repositories) InsertNewJurusan(jurusan []string) ([]uint, error) {
	db := repo.db

	AllJurusan, err := repo.GetAllJurusan()
	if err != nil {
		return []uint{}, err
	}

	// Have to insert jurusan that is currently not existed in db
	var result []uint

	var uninsertedJurusan []string
	found := false
	for _, h := range jurusan {
		for _, ah := range AllJurusan {
			if ah.Nama == h {
				found = true
				result = append(result, ah.Id)
			}
		}
		if !found {
			uninsertedJurusan = append(uninsertedJurusan, h)
		}
		found = false
	}

	// then insert uninsertedJurusan with transaction
	tx, _ := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	for _, u := range uninsertedJurusan {
		res, err := tx.Exec(
			"INSERT INTO jurusan (nama) VALUES (?)",
			u)
		if err != nil {
			tx.Rollback()
			return []uint{}, err
		}
		id, _ := res.LastInsertId()
		result = append(result, uint(id))
	}

	tx.Commit()
	// get all jurusan ids
	return result, nil
}

func (repo *repositories) InserNewMahasiswaJurusan(mahasiswaId uint, jurusanIds []uint) error {
	db := repo.db

	tx, _ := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	for _, jid := range jurusanIds {
		_, err := tx.Exec(
			"INSERT INTO mahasiswa_jurusan (id_mahasiswa, id_jurusan) VALUES (?,?)",
			mahasiswaId,
			jid,
		)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}
