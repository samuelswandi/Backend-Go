package repositories

import (
	"main/entities/responses"
)

func (repo *repositories) GetMahasiswaById(id uint) (responses.MahasiswaResponse, error) {
	db := repo.db

	query := `
	select m.id, m.nama, usia, gender, tanggal_registrasi, COALESCE(GROUP_CONCAT(DISTINCT (h.nama) SEPARATOR ", "), "None") as hobi,COALESCE(GROUP_CONCAT(DISTINCT (j.nama) SEPARATOR ", "), "None")as jurusan from mahasiswa m
    left join mahasiswa_hobi mh on mh.id_mahasiswa = m.id
    left join hobi h on h.id = mh.id_hobi
    left join mahasiswa_jurusan mj on mj.id_mahasiswa = m.id
    left join jurusan j on mj.id_jurusan = j.id
	where m.id = ?
	group by m.id;
	`

	rows, err := db.Query(query, id)
	if err != nil {
		return responses.MahasiswaResponse{}, err
	}
	defer rows.Close()

	var result responses.MahasiswaResponse
	for rows.Next() {
		var each = responses.MahasiswaResponse{}
		var err = rows.Scan(&each.Id, &each.Nama, &each.Usia, &each.Gender, &each.TanggalRegistrasi, &each.Hobi, &each.Jurusan)

		if err != nil {
			return responses.MahasiswaResponse{}, err
		}

		result = each
	}

	if err = rows.Err(); err != nil {
		return responses.MahasiswaResponse{}, err
	}

	return result, nil
}

func (repo *repositories) GetMahasiswaByName(name string) ([]responses.MahasiswaResponse, error) {
	db := repo.db

	query := `
	select m.id, m.nama, usia, gender, tanggal_registrasi, COALESCE(GROUP_CONCAT(DISTINCT (h.nama) SEPARATOR ", "), "None") as hobi,COALESCE(GROUP_CONCAT(DISTINCT (j.nama) SEPARATOR ", "), "None")as jurusan from mahasiswa m
    left join mahasiswa_hobi mh on mh.id_mahasiswa = m.id
    left join hobi h on h.id = mh.id_hobi
    left join mahasiswa_jurusan mj on mj.id_mahasiswa = m.id
    left join jurusan j on mj.id_jurusan = j.id
	where m.nama like ?
	group by m.id;
	`

	rows, err := db.Query(query, "%"+name+"%")
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
