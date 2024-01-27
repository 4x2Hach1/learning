package services

import (
	"context"

	"github.com/4x2Hach1/learning/next-go/api/gen/server"
)

type helloService struct {
	*serverInfr
}

func (s *helloService) Hello(ctx context.Context, p *server.HelloPayload) (string, error) {
	s.logger.Print("server.Hello")
	return string(p.Name), nil
}
