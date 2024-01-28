package services

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/4x2Hach1/learning/next-go/api/cache"
	"github.com/4x2Hach1/learning/next-go/api/gen/server"
)

var (
	slots   = make(chan bool, 3)
	slotsWg sync.WaitGroup
)

func heavy(update func(int) error) {
	slotsWg.Add(1)
	slots <- true

	for i := 0; i <= 100; i += 20 {
		err := update(i)
		fmt.Printf("heavy:%d, err:%v\n", i, err)
		time.Sleep(10 * time.Second)
	}

	slotsWg.Done()
	<-slots
}

type heavyService struct {
	*serverInfr
}

func (s *heavyService) CheckHeavy(ctx context.Context, p *server.CheckHeavyPayload) (int, error) {
	s.logger.Print("server.CheckHeavy")

	key := &cache.CacheHeavyRawKey{Hash: p.Key}
	val := &cache.CacheHeavyValue{}

	if err := s.cache.GetCache(key, val); err != nil {
		return 0, err
	}
	return val.Status, nil
}

func (s *heavyService) DeleteHeavy(ctx context.Context, p *server.DeleteHeavyPayload) (bool, error) {
	s.logger.Print("server.DeleteHeavy")

	key := &cache.CacheHeavyRawKey{Hash: p.Key}

	if err := s.cache.DeleteCache(key); err != nil {
		return false, err
	}
	return true, nil
}

func (s *heavyService) NewHeavy(ctx context.Context, p *server.NewHeavyPayload) (string, error) {
	s.logger.Print("server.NewHeavy")

	token := getUserFromCtx(ctx)
	key := &cache.CacheHeavyKey{User: token.UserId}
	val := &cache.CacheHeavyValue{Status: 0}

	go heavy(func(status int) error {
		val.Status = status
		if err := s.cache.SetCache(key, val, 3600); err != nil {
			return err
		}
		return nil
	})

	return key.Key(), nil
}
