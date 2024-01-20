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
	memories_column = []string{"id", "users_id", "title", "date", "location", "detail", "created_at", "updated_at"}
)

func TestMemories(t *testing.T) {
	db, mock, err := services.ExportSetUpMockDB()
	if err != nil {
		t.Fatal(err)
	}

	logger := log.New(os.Stderr, "[test] ", log.Ltime)
	srv := services.ExportNewMemoryService(db, logger)
	now := time.Now()

	tests := []struct {
		title string
		param *server.MemoriesPayload
		setup func(sqlmock.Sqlmock) int
	}{
		{
			"memories ok",
			&server.MemoriesPayload{Token: ""},
			func(s sqlmock.Sqlmock) int {
				s.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM memories`)).WillReturnRows(
					s.NewRows(memories_column).AddRow(1, 1, "Title", now, "https://maps.app.goo.gl/KWjr4KurSaRUVpC79", "Detail", now, now),
				)

				return 1
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			expected := tt.setup(mock)
			res, err := srv.Memories(context.Background(), tt.param)

			assert.Nil(t, err, "error must nil")
			assert.Equal(t, expected, len(res), "res must equal")
		})
	}
}
