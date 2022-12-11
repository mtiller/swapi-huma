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

type PersonDetails struct {
	Name      string `json:"name"`
	Gender    string `json:"gender"`
	Hair      string `json:"hair_color"`
	Height    string `json:"height"`
	Eyes      string `json:"eye_color"`
	Mass      string `json:"mass"`
	BirthYear string `json:"birth_year"`
}

func ListPeople(r *huma.Resource) {
	r.Get("list-people",
		"Get a paginated list of people",
		responses.OK().Model([]claxon.Emb[PersonDetails]{}).ContentType("text/plain"),
		responses.InternalServerError(),
	).Run(personListHandler)

}

func ShowPersonDetails(r *huma.Resource) {
	r.Get("person-details",
		"See details about a specific person",
		responses.OK().Model([]claxon.Emb[PersonDetails]{}).ContentType("text/plain"),
		responses.BadRequest(),
		responses.InternalServerError(),
	).Run(personHandler)
}

func personListHandler(ctx huma.Context) {
	database := GetService[*data.Database](ctx)

	people := lo.Map(database.People, getPersonDetails)

	ctx.Header().Set("Content-Type", "text/plain")
	WriteModel(http.StatusOK, "", ctx, people, nil)
}

func personHandler(ctx huma.Context, input struct {
	Id     int    `path:"id"`
	Accept string `header:"Accept"`
}) {
	database := GetService[*data.Database](ctx)
	log := middleware.GetLogger(ctx)

	id := 1
	selected, ok := lo.Find(database.People, func(person data.People) bool {
		return person.Id == id
	})

	if !ok {
		ctx.AddError(fmt.Errorf("No person found with id of %d", id))
	}

	if ctx.HasError() {
		ctx.WriteError(http.StatusBadRequest, "Unable to process request")
		return
	}

	details := getPersonDetails(selected, 0)
	rels := &details.Context

	rels.AddLink("homeworld", fmt.Sprintf("/planet/%d", selected.Fields.Homeworld), database.Planets[selected.Fields.Homeworld].Fields.Name)

	// TODO: Reverse relation to films
	for _, film := range database.Films {
		for _, c := range film.Fields.Characters {
			if c == selected.Id {
				rels.AddLink("film", fmt.Sprintf("/film/%d", film.Id), film.Fields.Title)
			}
		}
	}

	ctx.Header().Set("Content-Type", "text/plain")
	log.Infof("Accept: %s", input.Accept)
	WriteModel(http.StatusOK, input.Accept, ctx, details, nil)
}

func getPersonDetails(person data.People, index int) claxon.Emb[PersonDetails] {
	c := &claxon.Claxon{
		Self: fmt.Sprintf("/people/%d", person.Id),
	}
	c.AddLink("collection", "/people", "People")
	return claxon.Emb[PersonDetails]{
		Data: PersonDetails{
			Name:      person.Fields.Name,
			Gender:    person.Fields.Gender,
			Hair:      person.Fields.Hair,
			Height:    person.Fields.Height,
			Eyes:      person.Fields.Eyes,
			Mass:      person.Fields.Mass,
			BirthYear: person.Fields.BirthYear,
		},
		Context: *c,
	}
}

func personLink(person data.People, index int) claxon.Link {
	return claxon.Link{
		Rel:   "item",
		Title: person.Fields.Name,
		Href:  fmt.Sprintf("/person/%d", person.Id),
	}
}
