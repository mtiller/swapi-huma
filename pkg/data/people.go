package data

import (
	"encoding/json"
	"os"
	"time"
)

type People struct {
	Id     int    `json:"pk"`
	Model  string `json:"model"`
	Fields struct {
		Edited    string    `json:"edited"`
		Name      string    `json:"name"`
		Created   time.Time `json:"created"`
		Gender    string    `json:"gender"`
		Hair      string    `json:"hair_color"`
		Height    string    `json:"height"`
		Eyes      string    `json:"eye_color"`
		Mass      string    `json:"mass"`
		Homeworld int       `json:"homeworld"`
		BirthYear string    `json:"birth_year"`
	} `json:"fields"`
}

func LoadPeople() ([]People, error) {
	var ret []People
	bytes, err := os.ReadFile("./data/people.json")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, &ret)
	return ret, err
}
