package entities

import (
	"github.com/edgedb/edgedb-go"
)

type Dish struct {
	id          edgedb.OptionalUUID
	Title       string             `edgedb:"title" form:"title"`
	Duration    int16              `edgedb:"duration" form:"duration"`
	Description edgedb.OptionalStr `edgedb:"description" form:"description"`
	Comment     edgedb.OptionalStr `edgedb:"comment" form:"comment"`
	Ingredients []Ingredient       `edgedb:"ingredients"`
	Steps       []Step             `edgedb:"steps"`
}

func (dish Dish) GetEdgeName() string {
	return "Dish"
}

func (dish Dish) GetDeletproperty() string {
	return "title"
}

func (dish Dish) GetPropertyValue() string {
	return dish.Title
}

func (dish Dish) FillDefault() Dish {
	_, descExists := edgedb.OptionalStr.Get(dish.Description)
	if !descExists {
		desc := edgedb.NewOptionalStr("No description")
		dish.Description = desc
	}

	_, commExists := edgedb.OptionalStr.Get(dish.Comment)
	if !commExists {
		comment := edgedb.NewOptionalStr("No Comment")
		dish.Comment = comment
	}

	return dish
}
