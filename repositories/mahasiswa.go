package repositories

import (
	"main/entities"
	"main/entities/responses"
)

func (repo *repositories) GetAllMahasiswa() ([]responses.MahasiswaResponse, error) {
	db := repo.db

	query := `
	select m.id, m.nama, usia, gender, tanggal_registrasi, COALESCE(GROUP_CONCAT(DISTINCT (h.nama) SEPARATOR ", "), "None") as hobi,COALESCE(GROUP_CONCAT(DISTINCT (j.nama) SEPARATOR ", "), "None")as jurusan from mahasiswa m
    left join mahasiswa_hobi mh on mh.id_mahasiswa = m.id
    left join hobi h on h.id = mh.id_hobi
    left join mahasiswa_jurusan mj on mj.id_mahasiswa = m.id
    left join jurusan j on mj.id_jurusan = j.id
	group by m.id;
	`

	rows, err := db.Query(query)
	if err != nil {
		return []responses.MahasiswaResponse{}, err
	}
	defer rows.Close()

	var result []responses.MahasiswaResponse
	for rows.Next() {
		var each = responses.MahasiswaResponse{}
		var err = rows.Scan(&each.Id, &each.Nama, &each.Usia, &each.Gender, &each.TanggalRegistrasi, &each.Hobi, &each.Jurusan)

		if err != nil {
			return []responses.MahasiswaResponse{}, err
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		return []responses.MahasiswaResponse{}, err
	}

	return result, nil
}

func (repo *repositories) InsertNewMahasiswa(mahasiswa entities.Mahasiswa) (entities.Mahasiswa, error) {
	db := repo.db

	result, err := db.Exec(
		"INSERT INTO mahasiswa (nama, usia, gender, tanggal_registrasi) VALUES (?,?,?,?)",
		mahasiswa.Nama,
		mahasiswa.Usia,
		mahasiswa.Gender,
		mahasiswa.TanggalRegistrasi.Format("2006-01-02 15:04:05"),
	)
	if err != nil {
		return entities.Mahasiswa{}, err
	}

	newId, err := result.LastInsertId()
	if err != nil {
		return entities.Mahasiswa{}, err
	}

	return entities.Mahasiswa{
		Id:                uint(newId),
		Nama:              mahasiswa.Nama,
		Usia:              mahasiswa.Usia,
		Gender:            mahasiswa.Gender,
		TanggalRegistrasi: mahasiswa.TanggalRegistrasi,
	}, nil
}

