package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"connectrpc.com/grpcreflect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	connectcors "connectrpc.com/cors"
	"github.com/jmoiron/sqlx"
	"github.com/rs/cors"

	"github.com/4x2Hach1/learning/grpc-go/api/models"
	service_hellov1 "github.com/4x2Hach1/learning/grpc-go/api/services/hello/v1"
	service_userv1 "github.com/4x2Hach1/learning/grpc-go/api/services/user/v1"
)

// https://connectrpc.com/docs/go/getting-started/

func withCORS(h http.Handler) http.Handler {
	middleware := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: connectcors.AllowedMethods(),
		AllowedHeaders: connectcors.AllowedHeaders(),
		ExposedHeaders: connectcors.ExposedHeaders(),
	})
	return middleware.Handler(h)
}

func main() {
	mux := http.NewServeMux()

	// connect db
	DSN := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true&autocommit=0&sql_mode=%%27TRADITIONAL,NO_AUTO_VALUE_ON_ZERO,ONLY_FULL_GROUP_BY%%27",
		os.Getenv("DB_USER"), os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"),
	)
	sql, err := sqlx.Open("mysql", DSN)
	if err != nil {
		log.Fatal(err)
	}
	db := models.NewSqlDB(sql)

	// make groc handlers
	// hello
	path, handler := service_hellov1.NewHelloServer()
	mux.Handle(path, handler)
	// user
	path, handler = service_userv1.NewUserService(db)
	mux.Handle(path, handler)

	reflector := grpcreflect.NewStaticReflector(
		"proto.hello.v1.HelloService",
		"proto.user.v1.UserService",
	)
	mux.Handle((grpcreflect.NewHandlerV1(reflector)))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))

	log.Println("localhost:5555")
	http.ListenAndServe(
		"localhost:5555",
		// Use h2c so we can serve HTTP/2 without TLS.
		withCORS(h2c.NewHandler(mux, &http2.Server{})),
	)
}
