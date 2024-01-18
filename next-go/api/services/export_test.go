package services

import (
	"log"

	"github.com/4x2Hach1/learning/next-go/api/database/models"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

func ExportSetUpMockDB() (*models.Sql, sqlmock.Sqlmock, error) {
	driver, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	db := sqlx.NewDb(driver, "sqlmock")
	sql := models.NewSqlDB(db)

	return sql, mock, nil
}

func ExportNewAuthService(db *models.Sql, logger *log.Logger) authService {
	return authService{db, logger}
}

func ExportNewHelloService(db *models.Sql, logger *log.Logger) helloService {
	return helloService{db, logger}
}

func ExportNewUserService(db *models.Sql, logger *log.Logger) userService {
	return userService{db, logger}
}
