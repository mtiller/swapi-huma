package handlers

import "github.com/mtiller/go-claxon"

type Emb[T any] struct {
	Data    T
	Context claxon.Claxon
}

func (e Emb[T]) MarshalJSON() ([]byte, error) {
	return Marshal(e.Data, e.Context)
}