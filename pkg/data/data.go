package data

type Database struct {
	Films      []Film
	People     []People
	Plants     []Planet
	Species    []Species
	Starships  []Starship
	Transports []Transport
	Vehicles   []Vehicle
}

func Load() (Database, error) {
	films, err := LoadFilms()
	if err != nil {
		return Database{}, err
	}

	people, err := LoadPeople()
	if err != nil {
		return Database{}, err
	}

	planets, err := LoadPlanets()
	if err != nil {
		return Database{}, err
	}

	species, err := LoadSpecies()
	if err != nil {
		return Database{}, err
	}

	starships, err := LoadStarships()
	if err != nil {
		return Database{}, err
	}

	transports, err := LoadTransport()
	if err != nil {
		return Database{}, err
	}

	vehicles, err := LoadVehicles()
	if err != nil {
		return Database{}, err
	}

	return Database{
		Films:      films,
		People:     people,
		Plants:     planets,
		Species:    species,
		Starships:  starships,
		Transports: transports,
		Vehicles:   vehicles,
	}, nil
}
