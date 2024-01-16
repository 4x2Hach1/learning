package services

import (
	"log"

	server "github.com/4x2Hach1/learning/next-go/api/gen/server"
)

// server service example implementation.
// The example methods log the requests and return zero values.
type serversrvc struct {
	*helloService
	*userService
}

// NewServer returns the server service implementation.
func NewServer(logger *log.Logger) server.Service {
	return &serversrvc{
		&helloService{logger},
		&userService{logger},
	}
}
