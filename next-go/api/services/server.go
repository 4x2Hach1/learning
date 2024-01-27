package services

import (
	"log"

	"github.com/4x2Hach1/learning/next-go/api/cache"
	"github.com/4x2Hach1/learning/next-go/api/database/models"
	server "github.com/4x2Hach1/learning/next-go/api/gen/server"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type serverInfr struct {
	db     *models.Sql
	cache  *cache.Cache
	logger *log.Logger
}

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
	infr := &serverInfr{sql, c, logger}

	return &serversrvc{
		&authService{infr},
		&helloService{infr},
		&userService{infr},
		&memoryService{infr},
		&heavyService{infr},
	}
}
