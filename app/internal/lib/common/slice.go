package libCommon

import "slices"

type (
	ISliceManipulator[T any] interface {
		AddElement(e ...T)
	}
)

type (
	Slice[T any] []T
)

func (this *Slice[T]) AddElement(e ...T) {

	switch {
	case *this == nil:
		*this = e
	default:
		*this = slices.Concat(*this, e)
	}
}
