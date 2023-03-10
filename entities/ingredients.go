package entities

import "github.com/edgedb/edgedb-go"

type Ingredient struct {
	Name     string               `edgedb:"name" form:"name"`
	Comment  edgedb.OptionalStr   `edgedb:"comment" form:"comment"`
	Quantity edgedb.OptionalInt64 `edgedb:"quantity" form:"quantity"`
}
