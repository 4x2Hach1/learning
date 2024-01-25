package services

import (
	"log"

	"github.com/4x2Hach1/learning/next-go/api/cache"
	"github.com/4x2Hach1/learning/next-go/api/database/models"
	server "github.com/4x2Hach1/learning/next-go/api/gen/server"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

// server service example implementation.
// The example methods log the requests and return zero values.
type serversrvc struct {
	*authService
	*helloService
	*userService
	*memoryService
	*heavyService
}

// NewServer returns the server service implementation.
func NewServer(
	db *sqlx.DB,
	rds *redis.Client,
	logger *log.Logger,
) server.Service {
	sql := models.NewSqlDB(db)
	c := cache.NewCache(rds)

	return &serversrvc{
		&authService{sql, c, logger},
		&helloService{sql, c, logger},
		&userService{sql, c, logger},
		&memoryService{sql, c, logger},
		&heavyService{sql, c, logger},
	}
}
