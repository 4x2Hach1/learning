package hellov1

import (
	"context"
	"log"
	"net/http"

	"connectrpc.com/connect"

	hellov1 "github.com/4x2Hach1/learning/grpc-go/api/gen/proto/hello/v1"
	"github.com/4x2Hach1/learning/grpc-go/api/gen/proto/hello/v1/hellov1connect"
)

type HelloService struct{}

func NewHelloServer() (string, http.Handler) {
	srv := &HelloService{}
	path, handler := hellov1connect.NewHelloServiceHandler(srv)

	return path, handler
}

func (s *HelloService) Hello(
	ctx context.Context,
	req *connect.Request[hellov1.HelloRequest],
) (*connect.Response[hellov1.HelloResponse], error) {
	log.Println("HelloService Hello headers: ", req.Msg.Name)

	res := connect.NewResponse(&hellov1.HelloResponse{
		Message: "Hello " + req.Msg.Name,
	})
	res.Header().Set("Hello-Version", "v1")

	return res, nil
}
