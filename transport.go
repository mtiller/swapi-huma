package main

import (
	"encoding/json"
	"os"
	"time"
)

type Transport struct {
	Id     int    `json:"pk"`
	Model  string `json:"model"`
	Fields struct {
		Edited              string    `json:"edited"`
		Consumables         string    `json:"consumables"`
		Name                string    `json:"name"`
		Created             time.Time `json:"created"`
		Capacity            string    `json:"cargo_capacity"`
		Passengers          string    `json:"passengers"`
		MaxAtmosphericSpeed string    `json:"max_atmospheric_speed"`
		Crew                string    `json:"crew"`
		Length              string    `json:"length"`
		Model               string    `json:"model"`
		Cost                string    `json:"cost_in_credits"`
		Manufacturer        string    `json:"manufacturer"`
	} `json:"fields"`
}

func LoadTransport() ([]Transport, error) {
	var ret []Transport
	bytes, err := os.ReadFile("./data/transport.json")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, &ret)
	return ret, err
}
