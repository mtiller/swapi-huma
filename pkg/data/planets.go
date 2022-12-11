package data

import (
	"encoding/json"
	"os"
	"time"
)

type Planet struct {
	Id     int    `json:"pk"`
	Model  string `json:"model"`
	Fields struct {
		Edited        string    `json:"edited"`
		Climate       string    `json:"climate"`
		SurfaceWater  string    `json:"surface_water"`
		Name          string    `json:"name"`
		Diameter      string    `json:"diameter"`
		Created       time.Time `json:"created"`
		Terrain       string    `json:"terrain"`
		Gravity       string    `json:"gravity"`
		OrbitalPeriod string    `json:"orbital_period"`
		Population    string    `json:"population"`
	} `json:"fields"`
}

func LoadPlanets() ([]Planet, error) {
	var ret []Planet
	bytes, err := os.ReadFile("./data/planets.json")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, &ret)
	return ret, err
}
