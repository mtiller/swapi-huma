package handlers

import (
	"github.com/danielgtaylor/huma"
	"github.com/danielgtaylor/huma/responses"
)

func ListFilms(r *huma.Resource) {
	r.Get("list-films",
		"Get a paginated list of films",
		responses.OK().ContentType("text/plain"),
	).Run(filmHandler)
}

func filmHandler(ctx huma.Context) {
	ctx.Header().Set("Content-Type", "text/plain")
	ctx.Write([]byte(`{}`))
}
