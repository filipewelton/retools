package helpers

import "backend/internal/application/typings"

type Either[T any] struct {
	IsLeft        bool
	IsRight       bool
	LeftResponse  typings.Error
	RightResponse T
}

func (e *Either[T]) Left(err typings.Error) {
	e.IsLeft = true
	e.IsRight = false
	e.LeftResponse = err
}

func (e *Either[T]) Right(value T) {
	e.IsLeft = false
	e.IsRight = true
	e.RightResponse = value
}
