package main

import (
	"encoding/json"
	"os"
	"time"
)

type Species struct {
	Id     int    `json:"pk"`
	Model  string `json:"model"`
	Fields struct {
		Edited          string    `json:"edited"`
		Classification  string    `json:"classification"`
		Name            string    `json:"name"`
		Designation     string    `json:"designation"`
		Created         time.Time `json:"created"`
		EyeColors       string    `json:"eye_colors"`
		People          []int     `json:"people"`
		SkinColors      string    `json:"skin_colors"`
		Language        string    `json:"language"`
		HairColors      string    `json:"hair_colors"`
		Homeworld       int       `json:"homeworld"`
		AverageLifespan string    `json:"average_lifespan"`
		AverageHeight   string    `json:"average_height"`
	} `json:"fields"`
}

func LoadSpecies() ([]Species, error) {
	var ret []Species
	bytes, err := os.ReadFile("./data/species.json")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, &ret)
	return ret, err
}
