package app

import (
	"database/sql"
	"sclearn/learn-golang/helper"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:secret123@tcp(localhost:3305)/learn_golang")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}