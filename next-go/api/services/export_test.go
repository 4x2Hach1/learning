package services

import (
	"context"
	"log"
	"time"

	"github.com/4x2Hach1/learning/next-go/api/cache"
	"github.com/4x2Hach1/learning/next-go/api/database/models"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-redis/redismock/v9"
	"github.com/jmoiron/sqlx"
)

func ExportNewAuthService(db *models.Sql, cache *cache.Cache, logger *log.Logger) authService {
	return authService{&serverInfr{db, cache, logger}}
}

func ExportNewHelloService(db *models.Sql, cache *cache.Cache, logger *log.Logger) helloService {
	return helloService{&serverInfr{db, cache, logger}}
}

func ExportNewUserService(db *models.Sql, cache *cache.Cache, logger *log.Logger) userService {
	return userService{&serverInfr{db, cache, logger}}
}

func ExportNewMemoryService(db *models.Sql, cache *cache.Cache, logger *log.Logger) memoryService {
	return memoryService{&serverInfr{db, cache, logger}}
}

func ExportSetUpMockDB() (*models.Sql, sqlmock.Sqlmock, error) {
	driver, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	db := sqlx.NewDb(driver, "sqlmock")
	sql := models.NewSqlDB(db)

	return sql, mock, nil
}

func ExportSetUpCache() (*cache.Cache, redismock.ClientMock) {
	rds, mock := redismock.NewClientMock()
	c := cache.NewCache(rds)
	return c, mock
}

func ExportMakeToken(id int) context.Context {
	token := &UserJwtClaims{
		UserId: id,
		Exp:    int(time.Now().Add(time.Hour * 24).Unix()),
	}
	return context.WithValue(context.Background(), UserJwtClaims{}, token)
}
