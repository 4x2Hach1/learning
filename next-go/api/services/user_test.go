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
	users_column = []string{"id", "name", "email", "created_at", "updated_at"}
)

func TestUsers(t *testing.T) {
	db, mock, err := services.ExportSetUpMockDB()
	if err != nil {
		t.Fatal(err)
	}

	logger := log.New(os.Stderr, "[api] ", log.Ltime)
	srv := services.ExportNewUserService(db, logger)
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
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM users`)).WillReturnRows(
					s.NewRows(users_column).AddRow(1, "user", "test@example.com", now, now),
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
