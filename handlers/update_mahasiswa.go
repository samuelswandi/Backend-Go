package handlers

import (
	"main/entities"
	"main/entities/requests"
	"main/entities/responses"
	"main/lib"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *handlers) UpdateMahasiswa(c echo.Context) error {
	/*
		STEP:
		1. Check if user id exist
		2. If exist, proceed to update
	*/

	repo := h.repo
	response := responses.GenericResponse[entities.Mahasiswa]{}

	var request requests.UpdateMahasiswaRequest
	if err := c.Bind(&request); err != nil {
		response.Message = lib.GenericError(err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}
	if err := c.Validate(request); err != nil {
		response.Message = lib.GenericError(err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	// step 1
	res, err := repo.GetMahasiswaById(uint(request.Id))
	if err != nil {
		response.Message = lib.GenericError(err.Error())
		return c.JSON(http.StatusInternalServerError, response)
	}
	if res.Nama == "" {
		response.Message = lib.GenericError("id does not exist")
		return c.JSON(http.StatusBadRequest, response)
	}

	// step 2
	err = repo.UpdateMahasiswa(entities.Mahasiswa{
		Id:     request.Id,
		Nama:   request.Nama,
		Gender: uint(request.Gender),
		Usia:   uint(request.Usia),
	})
	if err != nil {
		response.Message = lib.GenericError(err.Error())
		return c.JSON(http.StatusInternalServerError, response)
	}

	response.Message = lib.SUCCESS
	response.Data = entities.Mahasiswa{
		Id:     request.Id,
		Nama:   request.Nama,
		Gender: uint(request.Gender),
		Usia:   uint(request.Usia),
	}
	return c.JSON(http.StatusCreated, response)
}
