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
	cache, _ := services.ExportSetUpCache()

	logger := log.New(os.Stderr, "[test] ", log.Ltime)
	srv := services.ExportNewMemoryService(db, cache, logger)
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

func TestMemoryByID(t *testing.T) {
	db, mock, err := services.ExportSetUpMockDB()
	if err != nil {
		t.Fatal(err)
	}
	cache, _ := services.ExportSetUpCache()

	logger := log.New(os.Stderr, "[test] ", log.Ltime)
	srv := services.ExportNewMemoryService(db, cache, logger)
	now := time.Now()

	tests := []struct {
		title string
		param *server.MemoryByIDPayload
		setup func(sqlmock.Sqlmock)
	}{
		{
			"memory ok",
			&server.MemoryByIDPayload{Token: "", ID: 1},
			func(s sqlmock.Sqlmock) {
				s.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM memories`)).WillReturnRows(
					s.NewRows(memories_column).AddRow(1, 1, "Title", now, "https://maps.app.goo.gl/KWjr4KurSaRUVpC79", "Detail", now, now),
				)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			tt.setup(mock)
			res, err := srv.MemoryByID(context.Background(), tt.param)

			assert.Nil(t, err, "error must nil")
			assert.NotNil(t, res, "res no nil")
		})
	}
}

func TestNewMemory(t *testing.T) {
	db, mock, err := services.ExportSetUpMockDB()
	if err != nil {
		t.Fatal(err)
	}
	cache, _ := services.ExportSetUpCache()

	logger := log.New(os.Stderr, "[test] ", log.Ltime)
	srv := services.ExportNewMemoryService(db, cache, logger)
	ctx := services.ExportMakeToken(1)
	now := time.Now()

	tests := []struct {
		title string
		param *server.NewMemoryPayload
		setup func(sqlmock.Sqlmock) bool
	}{
		{
			"insert memory ok",
			&server.NewMemoryPayload{
				Token: "", Title: "Title", Date: "2024-01-01",
				Location: "https://maps.app.goo.gl/KWjr4KurSaRUVpC79", Detail: "Detail",
			},
			func(s sqlmock.Sqlmock) bool {
				s.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM users`)).WillReturnRows(
					s.NewRows(users_column).AddRow(1, "user", "test@example.com", "password", now, now),
				)

				s.ExpectBegin()
				s.ExpectExec(regexp.QuoteMeta(`INSERT INTO memories`)).WillReturnResult(
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
			res, err := srv.NewMemory(ctx, tt.param)

			assert.Nil(t, err, "error must nil")
			assert.Equal(t, expected, res, "res must equal")
		})
	}
}

func TestDeleteMemory(t *testing.T) {
	db, mock, err := services.ExportSetUpMockDB()
	if err != nil {
		t.Fatal(err)
	}
	cache, _ := services.ExportSetUpCache()

	logger := log.New(os.Stderr, "[test] ", log.Ltime)
	srv := services.ExportNewMemoryService(db, cache, logger)
	ctx := services.ExportMakeToken(1)
	now := time.Now()

	tests := []struct {
		title string
		param *server.DeleteMemoryPayload
		setup func(sqlmock.Sqlmock) bool
	}{
		{
			"delete memory ok",
			&server.DeleteMemoryPayload{Token: "", ID: 1},
			func(s sqlmock.Sqlmock) bool {
				s.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM users`)).WillReturnRows(
					s.NewRows(users_column).AddRow(1, "user", "test@example.com", "password", now, now),
				)

				s.ExpectBegin()
				s.ExpectExec(regexp.QuoteMeta(`DELETE FROM memories`)).WillReturnResult(
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
			res, err := srv.DeleteMemory(ctx, tt.param)

			assert.Nil(t, err, "error must nil")
			assert.Equal(t, expected, res, "res must equal")
		})
	}
}

func TestUpdateMemory(t *testing.T) {
	db, mock, err := services.ExportSetUpMockDB()
	if err != nil {
		t.Fatal(err)
	}
	cache, _ := services.ExportSetUpCache()

	logger := log.New(os.Stderr, "[test] ", log.Ltime)
	srv := services.ExportNewMemoryService(db, cache, logger)
	ctx := services.ExportMakeToken(1)
	now := time.Now()

	tests := []struct {
		title string
		param *server.UpdateMemoryPayload
		setup func(sqlmock.Sqlmock) bool
	}{
		{
			"delete memory ok",
			&server.UpdateMemoryPayload{
				Token: "", ID: 1, Title: "Title", Date: "2024-01-01",
				Location: "https://maps.app.goo.gl/KWjr4KurSaRUVpC79", Detail: "Detail",
			},
			func(s sqlmock.Sqlmock) bool {
				s.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM users`)).WillReturnRows(
					s.NewRows(users_column).AddRow(1, "user", "test@example.com", "password", now, now),
				)

				s.ExpectBegin()
				s.ExpectExec(regexp.QuoteMeta(`UPDATE memories`)).WillReturnResult(
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
			res, err := srv.UpdateMemory(ctx, tt.param)

			assert.Nil(t, err, "error must nil")
			assert.Equal(t, expected, res, "res must equal")
		})
	}
}
