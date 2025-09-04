package annotation

import (
	"reflect"
)

type (
	annotation_t interface {
		prove()
	}
	Annotation struct{}
)

func (Annotation) prove() {}

func AssertAnnotation(target reflect.Value) {

	switch target.Interface().(type) {
	case annotation_t:
	default:
		panic("expect the type that embed the Annotation type")
	}
}
