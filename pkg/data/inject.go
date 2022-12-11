package data

import "github.com/samber/do"

func DatabaseService(i *do.Injector) (*Database, error) {
	database, err := Load()
	return &database, err
}
