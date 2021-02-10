package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Driver to conect to mysql
)

// Conection is used to conect to database
func Conection() (*sql.DB, error) {
	conection := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", conection)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	return db, nil
}
