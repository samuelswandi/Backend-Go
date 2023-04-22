package handlers

import (
	"main/entities/responses"
	"main/lib"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h *handlers) DeleteMahasiswa(c echo.Context) error {
	/*
		STEP:
		1. Check for id in query param and make sure it is integer
		2. Check if user id existed
		3. If exist, proceed to delete
	*/

	repo := h.repo
	var response responses.GenericResponse[string]

	// step 1
	id := c.QueryParam("id")
	if id == "" {
		response.Message = lib.GenericError("id not exist in param")
		return c.JSON(http.StatusBadRequest, response)
	}
	parsedId, err := strconv.Atoi(id)
	if err != nil {
		response.Message = lib.GenericError(err.Error())
		return c.JSON(http.StatusInternalServerError, response)
	}

	// step 2
	res, err := repo.GetMahasiswaById(uint(parsedId))
	if err != nil {
		response.Message = lib.GenericError(err.Error())
		return c.JSON(http.StatusInternalServerError, response)
	}
	if res.Nama == "" {
		response.Message = lib.GenericError("id does not exist")
		return c.JSON(http.StatusBadRequest, response)
	}

	// step 3
	if err := repo.DeleteMahasiswaById(uint(parsedId)); err != nil {
		response.Message = lib.GenericError(err.Error())
		return c.JSON(http.StatusInternalServerError, response)
	}

	response.Message = lib.SUCCESS
	return c.JSON(http.StatusCreated, response)
}
