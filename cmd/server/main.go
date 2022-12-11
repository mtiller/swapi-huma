package main

import (
	"github.com/mtiller/swapi/pkg/data"
	"github.com/samber/do"
)

func main() {

	injector := do.New()

	// provides CarService
	do.Provide(injector, data.DatabaseService)
}
