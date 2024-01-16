package services

import (
	"context"
	"log"

	"github.com/4x2Hach1/learning/next-go/api/gen/server"
)

type helloService struct {
	logger *log.Logger
}

func (s *helloService) Hello(ctx context.Context, p *server.HelloPayload) (res string, err error) {
	s.logger.Print("server.Hello")
	return string(p.Name), nil
}
