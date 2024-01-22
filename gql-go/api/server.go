package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/4x2Hach1/learning/gql-go/api/execs"
	"github.com/4x2Hach1/learning/gql-go/api/middlewares/complexity"
	"github.com/4x2Hach1/learning/gql-go/api/middlewares/directives"
	"github.com/4x2Hach1/learning/gql-go/api/middlewares/loaders"
	"github.com/4x2Hach1/learning/gql-go/api/middlewares/loggers"
	"github.com/4x2Hach1/learning/gql-go/api/resolvers"
	"github.com/4x2Hach1/learning/gql-go/api/services"
	"github.com/4x2Hach1/learning/gql-go/api/utils"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/rs/cors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"go.uber.org/zap"
)

func main() {
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{
			"*",
		},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	if os.Getenv("DEBUG") == "true" {
		boil.DebugMode = true
		debugWriter := zap.NewStdLog(utils.Logger())
		boil.DebugWriter = debugWriter.Writer()
	}

	db, err := utils.SetupDB()
	if err != nil {
		utils.LoggingFatal()("system-log", "server", err)
	}
	defer db.Close()

	services := services.New(db)
	server := handler.New(execs.NewExecutableSchema(execs.Config{
		Resolvers: &resolvers.Resolver{
			Srv:     services,
			Loaders: loaders.NewLoaders(services),
		},
		Directives: directives.Directive,
		Complexity: complexity.ComplexityConfig(),
	}))
	server.AddTransport(transport.Options{})
	server.AddTransport(transport.POST{})

	server.SetErrorPresenter(loggers.SetErrorPresenter)
	server.SetRecoverFunc(loggers.SetRecoverFunc)

	server.Use(extension.FixedComplexityLimit(100))
	// Develop
	if os.Getenv("DEBUG") == "true" {
		server.Use(extension.Introspection{})
	}

	// graphql
	http.Handle("/query",
		http.TimeoutHandler(
			cors.Handler(directives.AuthAndLoggerMiddleware(server)),
			90*time.Second,
			"timeout!",
		))

	// helthcheck
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "OK")
	})

	if os.Getenv("DEBUG") == "true" {
		http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	}

	utils.LoggingInfo()("system-log", "server", "GraphQL server start!")
	utils.LoggingFatal()("system-log", "server", http.ListenAndServe(":"+"3000", nil))
}
