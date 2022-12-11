package main

import (
	"github.com/mtiller/swapi/pkg/data"
	"github.com/samber/do"
)

func dependencies() *do.Injector {
	// Create DI container
	injector := do.New()

	// Add database to container
	do.Provide(injector, data.DatabaseService)

	return injector
}
