package api

import (
	"context"
	"log"

	server "github.com/4x2Hach1/learning/next-go/api/gen/server"
)

// server service example implementation.
// The example methods log the requests and return zero values.
type serversrvc struct {
	logger *log.Logger
}

// NewServer returns the server service implementation.
func NewServer(logger *log.Logger) server.Service {
	return &serversrvc{logger}
}

// Hello implements hello.
func (s *serversrvc) Hello(ctx context.Context, p *server.HelloPayload) (res string, err error) {
	s.logger.Print("server.Hello")
	return string(p.Name), nil
}
