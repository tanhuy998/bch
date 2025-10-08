package object

import "app/internal/objectPool"

type (
	ObjectRecycler[T any] struct {
		objectPool.ObjectRecycler[T]
	}
)
