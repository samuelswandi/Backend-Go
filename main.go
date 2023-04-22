package main

import (
	"fmt"
	"log"
	"main/handlers"
	"main/lib"
	"main/repositories"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// validator
type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func main() {
	e := echo.New()

	/* Middlewares */
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Validator = &CustomValidator{validator: validator.New()}
	// e.Use(middlewares.Cors()) if cors needed

	/* Database */
	db, err := lib.Connect()
	if err != nil {
		log.Println(err.Error())
		return
	}
	fmt.Println("[Database] Connected")

	/* Repositories */
	// using separate repositories db for future improvement such as
	// microservices or unit testing
	repo := repositories.NewRepositories(db)

	/* Handlers */
	handler := handlers.NewHandlers(repo)

	/* Routes */
	// This can be seperated later
	// Can be extended with using middlewares for jwt etc
	e.GET("/mahasiswa", handler.GetAllMahasiswa)
	e.POST("/createMahasiswa", handler.CreateMahasiswa)
	e.POST("/updateMahasiswa", handler.UpdateMahasiswa)
	e.DELETE("/deleteMahasiswa", handler.DeleteMahasiswa)

	// Run server
	e.Logger.Fatal(e.Start(":8080"))
}
