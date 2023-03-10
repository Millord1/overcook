package entities

import "github.com/edgedb/edgedb-go"

type Step struct {
	Content string             `edgedb:"content" form:"content"`
	Comment edgedb.OptionalStr `edgedb:"comment" form:"comment"`
}
