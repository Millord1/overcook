package entities

import (
	"github.com/edgedb/edgedb-go"
)

type Ingredient struct {
	Id   edgedb.UUID `edgedb:"id"`
	Name string      `edgedb:"name" form:"name"`
	// Comment  edgedb.OptionalStr   `edgedb:"comment" form:"comment"`
	Quantity edgedb.OptionalInt32 `edgedb:"quantity" form:"quantity"`
	Unity    edgedb.OptionalStr   `edgedb:"unity" form:"unity"`
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

func (ing Ingredient) FillDefault() Ingredient {
	_, qttExists := ing.Quantity.Get()
	if !qttExists {
		ing.Quantity = edgedb.NewOptionalInt32(0)
	}

	_, unityExists := ing.Quantity.Get()
	if !unityExists {
		ing.Unity = edgedb.NewOptionalStr("nothing")
	}

	return ing
}
