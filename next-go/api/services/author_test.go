package services_test

import (
	"context"
	"errors"
	"log"
	"os"
	"regexp"
	"testing"
	"time"

	"github.com/4x2Hach1/learning/next-go/api/gen/server"
	"github.com/4x2Hach1/learning/next-go/api/services"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestJWTAuth(t *testing.T) {
	db, mock, err := services.ExportSetUpMockDB()
	if err != nil {
		t.Fatal(err)
	}

	logger := log.New(os.Stderr, "[test] ", log.Ltime)
	srv := services.ExportNewAuthService(db, logger)

	tests := []struct {
		title string
		param string
		setup func(sqlmock.Sqlmock) error
	}{
		{
			"test auth ok",
			os.Getenv("TEST_TOKEN"),
			func(s sqlmock.Sqlmock) error {
				return nil
			},
		},
		{
			"test auth ng",
			"",
			func(s sqlmock.Sqlmock) error {
				return errors.New("invalid token")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			expected := tt.setup(mock)
			_, err := srv.JWTAuth(context.Background(), tt.param, nil)

			assert.Equal(t, expected, err, "error must equal")
		})
	}
}

func TestLogin(t *testing.T) {
	db, mock, err := services.ExportSetUpMockDB()
	if err != nil {
		t.Fatal(err)
	}

	logger := log.New(os.Stderr, "[test] ", log.Ltime)
	srv := services.ExportNewAuthService(db, logger)
	now := time.Now()

	tests := []struct {
		title string
		param *server.LoginPayload
		setup func(sqlmock.Sqlmock) error
	}{
		{
			"test login ok",
			&server.LoginPayload{
				Email:    "test@example.com",
				Password: "password",
			},
			func(s sqlmock.Sqlmock) error {
				s.ExpectQuery(regexp.QuoteMeta(`SELECT id FROM users`)).WillReturnRows(
					s.NewRows(users_column).AddRow(1, "user", "test@example.com", "password", now, now),
				)

				return nil
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			expected := tt.setup(mock)
			_, err := srv.Login(context.Background(), tt.param)

			assert.Equal(t, expected, err, "error must equal")
		})
	}
}
