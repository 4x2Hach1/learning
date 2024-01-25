package services_test

import (
	"context"
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

var (
	users_column = []string{"id", "name", "email", "password", "created_at", "updated_at"}
)

func TestAuthUser(t *testing.T) {
	db, mock, err := services.ExportSetUpMockDB()
	if err != nil {
		t.Fatal(err)
	}
	cache, _ := services.ExportSetUpCache()

	logger := log.New(os.Stderr, "[test] ", log.Ltime)
	srv := services.ExportNewUserService(db, cache, logger)
	ctx := services.ExportMakeToken(1)
	now := time.Now()

	tests := []struct {
		title string
		param *server.AuthUserPayload
		setup func(sqlmock.Sqlmock)
	}{
		{
			"users ok",
			&server.AuthUserPayload{Token: "token"},
			func(s sqlmock.Sqlmock) {
				s.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM users`)).WillReturnRows(
					s.NewRows(users_column).AddRow(1, "user", "test@example.com", "password", now, now),
				)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			tt.setup(mock)
			res, err := srv.AuthUser(ctx, tt.param)

			assert.Nil(t, err, "error must nil")
			assert.NotNil(t, res, "res not nil")
		})
	}
}

func TestUsers(t *testing.T) {
	db, mock, err := services.ExportSetUpMockDB()
	if err != nil {
		t.Fatal(err)
	}
	cache, _ := services.ExportSetUpCache()

	logger := log.New(os.Stderr, "[test] ", log.Ltime)
	srv := services.ExportNewUserService(db, cache, logger)
	now := time.Now()

	tests := []struct {
		title string
		param *server.UsersPayload
		setup func(sqlmock.Sqlmock) int
	}{
		{
			"users ok",
			&server.UsersPayload{Token: "token"},
			func(s sqlmock.Sqlmock) int {
				s.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM users`)).WillReturnRows(
					s.NewRows(users_column).AddRow(1, "user", "test@example.com", "password", now, now),
				)

				return 1
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			expected := tt.setup(mock)
			res, err := srv.Users(context.Background(), tt.param)

			assert.Nil(t, err, "error must nil")
			assert.Equal(t, expected, len(res), "res must equal")
		})
	}
}

func TestUserByID(t *testing.T) {
	db, mock, err := services.ExportSetUpMockDB()
	if err != nil {
		t.Fatal(err)
	}
	cache, _ := services.ExportSetUpCache()

	logger := log.New(os.Stderr, "[test] ", log.Ltime)
	srv := services.ExportNewUserService(db, cache, logger)
	now := time.Now()

	tests := []struct {
		title string
		param *server.UserByIDPayload
		setup func(sqlmock.Sqlmock)
	}{
		{
			"user ok",
			&server.UserByIDPayload{Token: "", ID: 1},
			func(s sqlmock.Sqlmock) {
				s.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM users`)).WillReturnRows(
					s.NewRows(users_column).AddRow(1, "user", "test@example.com", "password", now, now),
				)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			tt.setup(mock)
			res, err := srv.UserByID(context.Background(), tt.param)

			assert.Nil(t, err, "error must nil")
			assert.NotNil(t, res, "res not nil")
		})
	}
}

func TestNewUser(t *testing.T) {
	db, mock, err := services.ExportSetUpMockDB()
	if err != nil {
		t.Fatal(err)
	}
	cache, _ := services.ExportSetUpCache()

	logger := log.New(os.Stderr, "[test] ", log.Ltime)
	srv := services.ExportNewUserService(db, cache, logger)

	tests := []struct {
		title string
		param *server.NewUserPayload
		setup func(sqlmock.Sqlmock) bool
	}{
		{
			"new user ok",
			&server.NewUserPayload{
				Name:     "user",
				Email:    "test@example.com",
				Password: "password",
			},
			func(s sqlmock.Sqlmock) bool {
				// s.ExpectQuery("SELECT count").WillReturnRows(
				// 	sqlmock.NewRows([]string{"count"}).AddRow(0),
				// )

				s.ExpectBegin()
				s.ExpectExec(regexp.QuoteMeta(`INSERT INTO users`)).WillReturnResult(
					sqlmock.NewResult(1, 1),
				)
				s.ExpectCommit()

				return true
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			expected := tt.setup(mock)
			res, err := srv.NewUser(context.Background(), tt.param)

			assert.Nil(t, err, "error must nil")
			assert.Equal(t, expected, res, "res must equal")
		})
	}
}

func TestUpdateUser(t *testing.T) {
	db, mock, err := services.ExportSetUpMockDB()
	if err != nil {
		t.Fatal(err)
	}
	cache, _ := services.ExportSetUpCache()

	logger := log.New(os.Stderr, "[test] ", log.Ltime)
	srv := services.ExportNewUserService(db, cache, logger)
	ctx := services.ExportMakeToken(1)

	tests := []struct {
		title string
		param server.UpdateUserPayload
		setup func(sqlmock.Sqlmock) bool
	}{
		{
			"update user ok",
			server.UpdateUserPayload{
				Token: "",
				Name:  "user",
				Email: "test@example.com",
			},
			func(s sqlmock.Sqlmock) bool {
				s.ExpectBegin()
				s.ExpectExec(regexp.QuoteMeta(`UPDATE users`)).WillReturnResult(
					sqlmock.NewResult(1, 1),
				)
				s.ExpectCommit()

				return true
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			expected := tt.setup(mock)
			res, err := srv.UpdateUser(ctx, &tt.param)

			assert.Nil(t, err, "error must nil")
			assert.Equal(t, expected, res, "res must equal")
		})
	}
}
