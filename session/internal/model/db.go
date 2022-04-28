package model

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dsnString = "root:12345@tcp(127.0.0.1:3306)/test"
)

var db *sql.DB

func Setup() {
	dsn := dsnString

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
	}

	if err := db.Ping(); err != nil {
		fmt.Println(err)
	}
}
