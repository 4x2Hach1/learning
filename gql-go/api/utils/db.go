package utils

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func SetupDB() (*sql.DB, error) {
	uri := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true",
		os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"),
	)

	db, err := sql.Open("mysql", uri)
	db.SetMaxIdleConns(1)
	db.SetMaxOpenConns(3)

	if err != nil {
		return nil, err
	} else {
		return db, nil
	}
}
