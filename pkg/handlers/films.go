package handlers

import (
	"github.com/danielgtaylor/huma"
	"github.com/danielgtaylor/huma/middleware"
	"github.com/danielgtaylor/huma/responses"
	"github.com/mtiller/swapi/pkg/data"
)

func ListFilms(r *huma.Resource) {
	r.Get("list-films",
		"Get a paginated list of films",
		responses.OK().ContentType("text/plain"),
	).Run(filmHandler)
}

func filmHandler(ctx huma.Context) {
	database := GetService[*data.Database](ctx)
	log := middleware.GetLogger(ctx)

	log.Infof("Found %d films in database", len(database.Films))
	ctx.Header().Set("Content-Type", "text/plain")
	ctx.Write([]byte(`{}`))
}
