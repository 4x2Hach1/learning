package services_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/4x2Hach1/learning/next-go/api/cache"
	"github.com/4x2Hach1/learning/next-go/api/gen/server"
	"github.com/4x2Hach1/learning/next-go/api/services"
	"github.com/go-redis/redismock/v9"
	"github.com/stretchr/testify/assert"
)

func TestCheckHeavy(t *testing.T) {
	db, _, err := services.ExportSetUpMockDB()
	if err != nil {
		t.Fatal(err)
	}
	rds, cmock := services.ExportSetUpCache()

	logger := log.New(os.Stderr, "[test] ", log.Ltime)
	srv := services.ExportNewHeavyService(db, rds, logger)
	key := &cache.CacheHeavyKey{User: 1}
	val := &cache.CacheHeavyValue{Status: 50}

	tests := []struct {
		title string
		param *server.CheckHeavyPayload
		setup func(redismock.ClientMock) int
	}{
		{
			"check heavy ok",
			&server.CheckHeavyPayload{Token: "Token", Key: os.Getenv("HEAVY_KEY")},
			func(c redismock.ClientMock) int {
				c.ExpectGet(key.Key()).SetVal(val.Value())

				return 50
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			expected := tt.setup(cmock)
			res, err := srv.CheckHeavy(context.Background(), tt.param)

			assert.Nil(t, err, "error must nil")
			assert.Equal(t, expected, res, "res must equal")
		})
	}
}

func TestDeleteHeavy(t *testing.T) {
	db, _, err := services.ExportSetUpMockDB()
	if err != nil {
		t.Fatal(err)
	}
	rds, cmock := services.ExportSetUpCache()

	logger := log.New(os.Stderr, "[test] ", log.Ltime)
	srv := services.ExportNewHeavyService(db, rds, logger)
	key := &cache.CacheHeavyKey{User: 1}

	tests := []struct {
		title string
		param *server.DeleteHeavyPayload
		setup func(redismock.ClientMock) bool
	}{
		{
			"delete heavy ok",
			&server.DeleteHeavyPayload{Token: "Token", Key: os.Getenv("HEAVY_KEY")},
			func(c redismock.ClientMock) bool {
				c.ExpectDel(key.Key()).RedisNil()
				return true
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			expected := tt.setup(cmock)
			res, err := srv.DeleteHeavy(context.Background(), tt.param)

			assert.Nil(t, err, "error must nil")
			assert.Equal(t, expected, res, "res must equal")
		})
	}
}

func TestNewHeavy(t *testing.T) {
	db, _, err := services.ExportSetUpMockDB()
	if err != nil {
		t.Fatal(err)
	}
	rds, cmock := services.ExportSetUpCache()

	logger := log.New(os.Stderr, "[test] ", log.Ltime)
	srv := services.ExportNewHeavyService(db, rds, logger)
	ctx := services.ExportMakeToken(1)
	key := &cache.CacheHeavyKey{User: 1}
	val := &cache.CacheHeavyValue{Status: 100}

	tests := []struct {
		title string
		param *server.NewHeavyPayload
		setup func(redismock.ClientMock)
	}{
		{
			"new ok",
			&server.NewHeavyPayload{Token: "Token"},
			func(c redismock.ClientMock) {
				c.Regexp().
					ExpectSet(key.Key(), val.Value(), time.Duration(60)*time.Second).
					SetVal(val.Value())
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			tt.setup(cmock)
			res, err := srv.NewHeavy(ctx, tt.param)
			fmt.Println(res)

			assert.Nil(t, err, "error must nil")
			assert.NotEmpty(t, res, "res not empty")
		})
	}
}
