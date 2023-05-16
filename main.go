package main

import (
	"example/api"
	"example/graph"
	"example/logging"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultGRPCPort = 8080
const defaultHTTPPort = 8089

func main() {
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	zap.L().Info("connect to GraphQL playground", logging.AutoField("port", defaultGRPCPort))

	r := gin.Default()
	r.GET(api.HealthCheckString, api.NewHealthCheckHandler().HealthCheck)
	r.GET("/user", api.NewUserHandler().GetUser)

	go func() {
		r.Run(":8089")
	}()
	if err := http.ListenAndServe(":"+strconv.Itoa(defaultGRPCPort), nil); err != nil {
		zap.L().Error("Server error", zap.Error(err))

	}
}
