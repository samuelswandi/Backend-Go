package lib

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

// Database config
func Connect() (*sql.DB, error) {
	config := mysql.Config{
		User:                 "root",
		Passwd:               "password",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "mahasiswa",
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		return nil, err
	}

	return db, nil
}

// Database for testing
func ConnectTesting() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/mahasiswa")
	if err != nil {
		return nil, err
	}

	return db, nil
}
