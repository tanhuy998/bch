package annotationScope

import (
	"fmt"
	"reflect"
)

type (
	Struct struct {
	}
)

func (Struct) _struct() {}

func AssertStructAnnotation(targetType reflect.Type) {

	switch {
	case !targetType.Implements(reflect.TypeFor[StructAnnotationConstraint]()):
		panic(
			fmt.Sprintf(
				"Could not apply %s that is not struct annotation",
				targetType.Name(),
			),
		)
	}
}
