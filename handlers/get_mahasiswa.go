package handlers

import (
	"main/entities/responses"
	"main/lib"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h *handlers) GetAllMahasiswa(c echo.Context) error {
	/*
		STEP:
		1. jika ada id, maka cari berdasarkan id
		2. jika ada nama cari berdasarkan nama
		3. jika tidak ada, maka cari semuanya
	*/

	repo := h.repo
	response := responses.GenericResponse[[]responses.MahasiswaResponse]{}

	// step 1
	id := c.QueryParam("id")
	if id != "" {
		parsedId, err := strconv.Atoi(id)
		if err != nil {
			response.Message = lib.GenericError(err.Error())
			return c.JSON(http.StatusBadRequest, response)
		}

		result, err := repo.GetMahasiswaById(uint(parsedId))
		if err != nil {
			response.Message = lib.GenericError(err.Error())
			return c.JSON(http.StatusInternalServerError, response)
		}
		if result.Nama == "" {
			response.Message = lib.GenericError("id does not exist")
			return c.JSON(http.StatusBadRequest, response)
		}

		response.Data = append(response.Data, result)
		response.Message = lib.SUCCESS
		return c.JSON(http.StatusOK, response)
	}

	// step 2
	nama := c.QueryParam("nama")
	if nama != "" {
		result, err := repo.GetMahasiswaByName(nama)
		if err != nil {
			response.Message = lib.GenericError(err.Error())
			return c.JSON(http.StatusInternalServerError, response)
		}

		response.Data = result
		response.Message = lib.SUCCESS
		return c.JSON(http.StatusOK, response)
	}

	// step 3
	result, err := repo.GetAllMahasiswa()
	if err != nil {
		response.Message = lib.GenericError(err.Error())
		return c.JSON(http.StatusInternalServerError, response)
	}

	response.Message = lib.SUCCESS
	response.Data = result
	return c.JSON(http.StatusOK, response)
}
