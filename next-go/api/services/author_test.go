package services_test

import (
	"context"
	"errors"
	"log"
	"os"
	"testing"

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
			"token",
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
