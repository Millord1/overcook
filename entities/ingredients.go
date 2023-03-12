package entities

import "github.com/edgedb/edgedb-go"

type Ingredient struct {
	Name     string               `edgedb:"name" form:"name"`
	Comment  edgedb.OptionalStr   `edgedb:"comment" form:"comment"`
	Quantity edgedb.OptionalInt64 `edgedb:"quantity" form:"quantity"`
}

func (ing Ingredient) GetEdgeName() string {
	return "Ingredient"
}

func (ing Ingredient) GetDeleteproperty() string {
	return "name"
}

func (ing Ingredient) GetPropertyValue() string {
	return ing.Name
}
