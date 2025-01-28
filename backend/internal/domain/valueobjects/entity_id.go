package valueobjects

import "github.com/google/uuid"

type EntityID struct {
	value string
}

func (e *EntityID) Generate() {
	str, err := uuid.NewV7()

	if err != nil {
		panic(err)
	}

	e.value = str.String()
}

func (e *EntityID) GetValue() string {
	return e.value
}

func (e *EntityID) Assign(value string) {
	e.value = value
}
