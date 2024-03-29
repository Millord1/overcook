package entities

import "github.com/edgedb/edgedb-go"

type Step struct {
	Content string             `edgedb:"content" form:"content"`
	Comment edgedb.OptionalStr `edgedb:"comment" form:"comment"`
}

func (step Step) GetEdgeName() string {
	return "Step"
}

func (step Step) GetDeleteproperty() string {
	return "content"
}

func (step Step) GetPropertyValue() string {
	return step.Content
}

func (step Step) FillDefault() Step {
	_, exists := edgedb.OptionalStr.Get(step.Comment)
	if !exists {
		step.Comment = edgedb.NewOptionalStr("No comment")
	}
	return step
}
