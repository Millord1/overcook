package entities

type EdgeEntity interface {
	GetEdgeName() string
	GetDeleteproperty() string
	GetPropertyValue() string
}

func GetDbName(e EdgeEntity) string {
	return e.GetEdgeName()
}

func GetProperty(e EdgeEntity) string {
	return e.GetDeleteproperty()
}

func GetValue(e EdgeEntity) string {
	return e.GetPropertyValue()
}
