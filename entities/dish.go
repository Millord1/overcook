package entities

import "github.com/edgedb/edgedb-go"

type Dish struct {
	id          edgedb.OptionalUUID
	Title       string             `edgedb:"title" form:"title"`
	Duration    int64              `edgedb:"duration" form:"duration"`
	Description edgedb.OptionalStr `edgedb:"description" form:"description"`
	Comment     edgedb.OptionalStr `edgedb:"comment" form:"comment"`
	Ingredients []Ingredient       `edgedb:"ingredients"`
	Steps       []Step             `edgedb:"steps"`
}
