package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/MatsuoTakuro/starwars/graph/generated"
	"github.com/MatsuoTakuro/starwars/graph/resolver"
)

const defaultPort = "8082"
const title = "gqlgen-starwars"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(resolver.NewResolver()))
	srv.AroundFields(func(ctx context.Context, next graphql.Resolver) (res interface{}, err error) {
		rc := graphql.GetFieldContext(ctx)
		fmt.Println("\nEntered", rc.Object, rc.Field.Name)
		res, err = next(ctx)
		fmt.Println("Output", rc.Object, rc.Field.Name, "=>", res, err)
		return res, err
	})

	http.Handle("/", playground.Handler(title, "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for %s", port, title)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
