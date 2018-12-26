package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/99designs/gqlgen/handler"
	"github.com/nvcnvn/graph-authorize/graph"
)

func main() {
	http.Handle("/", handler.Playground("Authorize", "/query"))

	cfg := graph.Config{
		Resolvers: &graph.Resolver{},
	}

	http.Handle("/query", handler.GraphQL(
		graph.NewExecutableSchema(cfg),
		handler.RecoverFunc(func(ctx context.Context, err interface{}) error {
			// send this panic somewhere
			log.Print(err)
			debug.PrintStack()
			return errors.New("user message on panic")
		}),
	))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
