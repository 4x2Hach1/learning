package services_test

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/4x2Hach1/learning/next-go/api/gen/server"
	"github.com/4x2Hach1/learning/next-go/api/services"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	db, mock, err := services.ExportSetUpMockDB()
	if err != nil {
		t.Fatal(err)
	}

	logger := log.New(os.Stderr, "[test] ", log.Ltime)
	srv := services.ExportNewHelloService(db, logger)

	tests := []struct {
		title string
		param *server.HelloPayload
		setup func(sqlmock.Sqlmock) string
	}{
		{
			"test hello",
			&server.HelloPayload{Name: "user"},
			func(s sqlmock.Sqlmock) string {
				return "user"
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			expected := tt.setup(mock)
			res, err := srv.Hello(context.Background(), tt.param)

			assert.Nil(t, err, "error must nil")
			assert.Equal(t, expected, res, "res must equal")
		})
	}
}
