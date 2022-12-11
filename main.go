package main

import "fmt"

func main() {
	films, err := LoadFilms()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Loaded %d films\n", len(films))

	people, err := LoadPeople()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Loaded %d people\n", len(people))

	planets, err := LoadPlanets()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Loaded %d planets\n", len(planets))

	species, err := LoadSpecies()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Loaded %d species\n", len(species))

	starships, err := LoadStarships()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Loaded %d starships\n", len(starships))

	transports, err := LoadTransport()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Loaded %d tranports\n", len(transports))

	vehicles, err := LoadVehicles()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Loaded %d vehicles\n", len(vehicles))

}
