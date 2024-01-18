package services

import (
	"log"

	"github.com/4x2Hach1/learning/next-go/api/database/models"
	server "github.com/4x2Hach1/learning/next-go/api/gen/server"
	"github.com/jmoiron/sqlx"
)

// server service example implementation.
// The example methods log the requests and return zero values.
type serversrvc struct {
	*authService
	*helloService
	*userService
}

// NewServer returns the server service implementation.
func NewServer(db *sqlx.DB, logger *log.Logger) server.Service {
	sql := models.NewSqlDB(db)
	return &serversrvc{
		&authService{sql, logger},
		&helloService{sql, logger},
		&userService{sql, logger},
	}
}
