package entities

import "github.com/edgedb/edgedb-go"

type EdgeResponse struct {
	Id edgedb.UUID `edgedb:"id"`
}
