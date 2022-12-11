package main

import (
	"encoding/json"
	"os"
)

type Starship struct {
	Id     int    `json:"pk"`
	Model  string `json:"model"`
	Fields struct {
		Pilots []int  `json:"pilots"`
		MGLT   string `json:"MGLT"`
		Class  string `json:"starship_class"`
		Rating string `json:"hyperdrive_rating"`
	} `json:"fields"`
}

func LoadStarships() ([]Starship, error) {
	var ret []Starship
	bytes, err := os.ReadFile("./data/starships.json")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, &ret)
	return ret, err
}
