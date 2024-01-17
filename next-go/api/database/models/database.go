package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Sql struct {
	db *sqlx.DB
}

func NewSqlDB(db *sqlx.DB) *Sql {
	return &Sql{db: db}
}
