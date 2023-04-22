package handlers

import (
	"main/entities"
	"main/entities/requests"
	"main/entities/responses"
	"main/lib"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func (h *handlers) CreateMahasiswa(c echo.Context) error {
	/*
		STEP:
		1. insert semua hobi
		2. insert semua jurusan
		3. insert mahasiswa baru
		4. insert mapping mahasiswa_hobi
		5. insert mapping mahasiswa_jurusan
	*/

	repo := h.repo
	var response responses.GenericResponse[entities.Mahasiswa]

	var request requests.CreateMahasiswaRequest
	if err := c.Bind(&request); err != nil {
		response.Message = lib.GenericError(err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}
	if err := c.Validate(request); err != nil {
		response.Message = lib.GenericError(err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	// step 1
	hobiIds, err := repo.InsertNewHobi(request.Hobi)
	if err != nil {
		response.Message = lib.GenericError(err.Error())
		return c.JSON(http.StatusInternalServerError, response)
	}

	// step 2
	jurusanIds, err := repo.InsertNewJurusan(request.Jurusan)
	if err != nil {
		response.Message = lib.GenericError(err.Error())
		return c.JSON(http.StatusInternalServerError, response)
	}

	// step 3
	newMahasiswa, err := repo.InsertNewMahasiswa(entities.Mahasiswa{
		Nama:              request.Nama,
		Usia:              request.Usia,
		Gender:            request.Gender,
		TanggalRegistrasi: time.Now(),
	})
	if err != nil {
		response.Message = lib.GenericError(err.Error())
		return c.JSON(http.StatusInternalServerError, response)
	}

	// step 4
	err = repo.InserNewMahasiswaHobi(newMahasiswa.Id, hobiIds)
	if err != nil {
		response.Message = lib.GenericError(err.Error())
		return c.JSON(http.StatusInternalServerError, response)
	}

	// step 5
	err = repo.InserNewMahasiswaJurusan(newMahasiswa.Id, jurusanIds)
	if err != nil {
		response.Message = lib.GenericError(err.Error())
		return c.JSON(http.StatusInternalServerError, response)
	}

	response.Message = lib.SUCCESS
	response.Data = newMahasiswa
	return c.JSON(http.StatusCreated, response)
}
