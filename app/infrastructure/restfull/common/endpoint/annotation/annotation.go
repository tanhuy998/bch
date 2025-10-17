package annotation

import (
	"fmt"
	"reflect"
)

type (
	IEndpointAnnotation interface {
		prove()
	}
	Annotation struct{}
)

func (Annotation) prove() {}

func AssertAnnotation(target reflect.Value) {

	if !TryAssertAnnotation(target) {

		panic(
			fmt.Sprintf(
				"%s is not valid annotation (not embed Annotation struct)",
				target.Type().Name(),
			),
		)
	}
}

func TryAssertAnnotation(target reflect.Value) bool {

	switch target.Interface().(type) {
	case IEndpointAnnotation:
		return true
	default:
		return false
	}
}
