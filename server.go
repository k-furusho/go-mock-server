package main

import (
	"go-mock-server/graph"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8081"

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = defaultPort
    }

    srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

    http.Handle("/", playground.Handler("GraphQL playground", "/query"))
    http.Handle("/query", srv)

    log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
    err := http.ListenAndServe(":"+port, nil)
    if err != nil {
        log.Fatalf("サーバー起動エラー: %v", err)
    }
}