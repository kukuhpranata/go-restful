package app

import (
	"database/sql"
	"kukuh/go-restful/helper"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:113322@tcp(localhost:3306)/go_restful")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
