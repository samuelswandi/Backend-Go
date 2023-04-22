package handlers

import (
	"main/repositories"

	"github.com/labstack/echo/v4"
)

type handlers struct {
	repo repositories.Repositories
}

type Handlers interface {
	CreateMahasiswa(c echo.Context) error
	UpdateMahasiswa(c echo.Context) error
	GetAllMahasiswa(c echo.Context) error
	DeleteMahasiswa(c echo.Context) error
}

func NewHandlers(repo repositories.Repositories) Handlers {
	return &handlers{repo: repo}
}
