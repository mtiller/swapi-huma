package handlers

import (
	"net/http"

	"github.com/danielgtaylor/huma"
	"github.com/danielgtaylor/huma/responses"
	"github.com/mtiller/swapi/pkg/data"
	"github.com/samber/lo"
)

func ListFilms(r *huma.Resource) {
	r.Get("list-films",
		"Get a paginated list of films",
		responses.OK().Model([]FilmDetails{}).ContentType("text/plain"),
	).Run(filmHandler)
}

type FilmDetails struct {
	Title    string `json:"title"`
	Episode  int    `json:"episode"`
	Director string `json:"directory"`
	Release  string `json:"released"`
}

func getFilmDetails(film data.Film, index int) FilmDetails {
	return FilmDetails{
		Title:    film.Fields.Title,
		Episode:  film.Fields.EpisodeId,
		Director: film.Fields.Director,
		Release:  film.Fields.Release,
	}
}

func filmHandler(ctx huma.Context) {
	database := GetService[*data.Database](ctx)

	films := lo.Map(database.Films, getFilmDetails)

	ctx.Header().Set("Content-Type", "text/plain")
	ctx.WriteModel(http.StatusOK, films)
}
