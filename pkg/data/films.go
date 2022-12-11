package data

import (
	"encoding/json"
	"os"
	"time"
)

type Film struct {
	Id     int    `json:"pk"`
	Model  string `json:"model"`
	Fields struct {
		Starships  []int     `json:"starships"`
		Edited     string    `json:"edited"`
		Vehicles   []int     `json:"vehicles"`
		Planets    []int     `json:"planets"`
		Title      string    `json:"title"`
		Created    time.Time `json:"created"`
		EpisodeId  int       `json:"episode_id"`
		Director   string    `json:"director"`
		Release    string    `json:"release_date"`
		Opening    string    `json:"string"`
		Characters []int     `json:"characters"`
		Species    []int     `json:"species"`
	} `json:"fields"`
}

func LoadFilms() ([]Film, error) {
	var ret []Film
	bytes, err := os.ReadFile("./data/films.json")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, &ret)
	return ret, err
}
