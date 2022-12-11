package data

import (
	"encoding/json"
	"os"
)

type Vehicle struct {
	Id     int    `json:"pk"`
	Model  string `json:"model"`
	Fields struct {
		Pilots []int  `json:"pilots"`
		Class  string `json:"vehicle_class"`
	} `json:"fields"`
}

func LoadVehicles() ([]Vehicle, error) {
	var ret []Vehicle
	bytes, err := os.ReadFile("./data/vehicles.json")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, &ret)
	return ret, err
}
