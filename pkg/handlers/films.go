package handlers

import (
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma"
	"github.com/danielgtaylor/huma/responses"
	"github.com/mtiller/go-claxon"
	"github.com/mtiller/swapi/pkg/data"
	"github.com/samber/lo"
)

type FilmDetails struct {
	Title    string `json:"title"`
	Episode  int    `json:"episode"`
	Director string `json:"directory"`
	Release  string `json:"released"`
	Self     string `json:"$self,omitempty"`
}

func ListFilms(r *huma.Resource) {
	r.Get("list-films",
		"Get a paginated list of films",
		responses.OK().Model([]FilmDetails{}).ContentType("text/plain"),
		responses.InternalServerError(),
	).Run(filmHandler)
}

func filmHandler(ctx huma.Context) {
	database := GetService[*data.Database](ctx)

	rels := claxon.Claxon{
		Links: lo.Map(database.Films, filmLink),
	}

	films := lo.Map(database.Films, getFilmDetails)

	ctx.Header().Set("Content-Type", "text/plain")
	WriteModel(http.StatusOK, ctx, films, rels)
}

func getFilmDetails(film data.Film, index int) FilmDetails {
	return FilmDetails{
		Title:    film.Fields.Title,
		Episode:  film.Fields.EpisodeId,
		Director: film.Fields.Director,
		Release:  film.Fields.Release,
		Self:     fmt.Sprintf("/film/%d", film.Id),
	}
}

func filmLink(film data.Film, index int) claxon.Link {
	return claxon.Link{
		Rel:   "item",
		Title: film.Fields.Title,
		Href:  fmt.Sprintf("/film/%d", film.Id),
	}
}
