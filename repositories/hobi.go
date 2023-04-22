package repositories

import (
	"main/entities"
)

func (repo *repositories) GetAllHobi() ([]entities.Hobi, error) {
	db := repo.db

	rows, err := db.Query("select * from hobi")
	if err != nil {
		return []entities.Hobi{}, err
	}
	defer rows.Close()

	var result []entities.Hobi

	for rows.Next() {
		var each = entities.Hobi{}
		var err = rows.Scan(&each.Id, &each.Nama)

		if err != nil {
			return []entities.Hobi{}, err
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		return []entities.Hobi{}, err
	}

	return result, nil
}

func (repo *repositories) InsertNewHobi(hobi []string) ([]uint, error) {
	db := repo.db

	AllHobi, err := repo.GetAllHobi()
	if err != nil {
		return []uint{}, err
	}

	// Have to insert hobi that is currently not existed in db
	var result []uint

	var uninsertedHobi []string
	found := false
	for _, h := range hobi {
		for _, ah := range AllHobi {
			if ah.Nama == h {
				found = true
				result = append(result, ah.Id)
			}
		}
		if !found {
			uninsertedHobi = append(uninsertedHobi, h)
		}
		found = false
	}

	// then insert uninsertedHobi with transaction
	tx, _ := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	for _, u := range uninsertedHobi {
		res, err := tx.Exec(
			"INSERT INTO hobi (nama) VALUES (?)",
			u)
		if err != nil {
			tx.Rollback()
			return []uint{}, err
		}
		id, _ := res.LastInsertId()
		result = append(result, uint(id))
	}

	tx.Commit()
	// get all hobi ids
	return result, nil
}


func (repo *repositories) InserNewMahasiswaHobi(mahasiswaId uint, hobiIds []uint) error {
	db := repo.db

	tx, _ := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	for _, hid := range hobiIds {
		_, err := tx.Exec(
			"INSERT INTO mahasiswa_hobi (id_mahasiswa, id_hobi) VALUES (?,?)",
			mahasiswaId,
			hid,
		)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}
