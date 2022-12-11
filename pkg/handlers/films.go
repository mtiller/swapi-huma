package handlers

import (
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma"
	"github.com/danielgtaylor/huma/middleware"
	"github.com/danielgtaylor/huma/responses"
	"github.com/mtiller/go-claxon"
	"github.com/mtiller/swapi/pkg/data"
	"github.com/samber/lo"
)

type FilmDetails struct {
	Title    string `json:"title"`
	Episode  int    `json:"episode"`
	Director string `json:"director"`
	Release  string `json:"released"`
}

func ListFilms(r *huma.Resource) {
	r.Get("list-films",
		"Get a paginated list of films",
		responses.OK().Model([]Emb[FilmDetails]{}).ContentType("text/plain"),
		responses.InternalServerError(),
	).Run(filmListHandler)

}

func ShowFilmDetails(r *huma.Resource) {
	r.Get("film-details",
		"See details about a specific film",
		responses.OK().Model([]Emb[FilmDetails]{}).ContentType("text/plain"),
		responses.BadRequest(),
		responses.InternalServerError(),
	).Run(filmHandler)
}

func filmListHandler(ctx huma.Context) {
	database := GetService[*data.Database](ctx)

	rels := claxon.Claxon{
		Links: lo.Map(database.Films, filmLink),
	}

	films := lo.Map(database.Films, getFilmDetails)

	ctx.Header().Set("Content-Type", "text/plain")
	WriteModel(http.StatusOK, "", ctx, films, rels)
}

func filmHandler(ctx huma.Context, input struct {
	Id     int    `path:"id"`
	Accept string `header:"Accept"`
}) {
	database := GetService[*data.Database](ctx)
	log := middleware.GetLogger(ctx)
	log.Infof("Input was %+v", input)

	rels := &claxon.Claxon{}
	rels.AddLink("collection", "/film", "Films")

	id := 1
	selected, ok := lo.Find(database.Films, func(film data.Film) bool {
		return film.Id == id
	})
	if !ok {
		ctx.AddError(fmt.Errorf("No film found with id of %d", id))
	}

	if ctx.HasError() {
		ctx.WriteError(http.StatusBadRequest, "Unable to process request")
		return
	}

	for _, character := range selected.Fields.Characters {
		rels.AddLink("character", fmt.Sprintf("/character/%d", character), database.People[character].Fields.Name)
	}

	for _, starship := range selected.Fields.Starships {
		rels.AddLink("starship", fmt.Sprintf("/starship/%d", starship), database.Starships[starship].Fields.Class)
	}

	for _, planet := range selected.Fields.Planets {
		rels.AddLink("planet", fmt.Sprintf("/planet/%d", planet), database.Planets[planet].Fields.Name)
	}

	for _, vehicle := range selected.Fields.Vehicles {
		rels.AddLink("vehicle", fmt.Sprintf("/vehicle/%d", vehicle), database.Vehicles[vehicle].Fields.Class)
	}

	ctx.Header().Set("Content-Type", "text/plain")
	log.Infof("Accept: %s", input.Accept)
	WriteModel(http.StatusOK, input.Accept, ctx, getFilmDetails(selected, 0), *rels)
}

func getFilmDetails(film data.Film, index int) Emb[FilmDetails] {
	c := &claxon.Claxon{}
	c.AddLink("self", fmt.Sprintf("/film/%d", film.Id))
	c.AddLink("collection", "/film")
	return Emb[FilmDetails]{
		Data: FilmDetails{
			Title:    film.Fields.Title,
			Episode:  film.Fields.EpisodeId,
			Director: film.Fields.Director,
			Release:  film.Fields.Release,
		},
		Context: *c,
	}
}

func filmLink(film data.Film, index int) claxon.Link {
	return claxon.Link{
		Rel:   "item",
		Title: film.Fields.Title,
		Href:  fmt.Sprintf("/film/%d", film.Id),
	}
}
