package main

import (
	"log"

	todo "github.com/omenstudio/fastgql/example/config"
	"github.com/omenstudio/fastgql/graphql/handler"
	"github.com/omenstudio/fastgql/graphql/playground"
	"github.com/valyala/fasthttp"
)

func main() {

	playground := playground.Handler("Todo", "/query")
	gqlHandler := handler.NewDefaultServer(todo.NewExecutableSchema(todo.New())).Handler()

	h := func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/query":
			gqlHandler(ctx)
		case "/":
			playground(ctx)
		default:
			ctx.Error("not found", fasthttp.StatusNotFound)
		}
	}

	log.Fatal(fasthttp.ListenAndServe(":8081", h))
}
