package main

import (
	"github.com/danielgtaylor/huma"
	"github.com/danielgtaylor/huma/cli"
	"github.com/danielgtaylor/huma/responses"
	"github.com/mtiller/swapi/pkg/handlers"
	"github.com/samber/do"
)

func application(injector *do.Injector) *cli.CLI {
	// Create a new router & CLI with default middleware.
	app := cli.NewRouter("Star Wars API", "0.2.0")

	app.Middleware(handlers.Inject(injector))

	// Declare the root resource and a GET operation on it.
	app.Resource("/").Get("get-root", "Get a short text message",
		// The only response is HTTP 200 with text/plain
		responses.OK().ContentType("text/plain"),
	).Run(func(ctx huma.Context) {
		// This is he handler function for the operation. Write the response.
		ctx.Header().Set("Content-Type", "text/plain")
		ctx.Write([]byte("Hello, world"))
	})

	handlers.ListFilms(app.Resource("/film"))

	return app
}
