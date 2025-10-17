package annotationScope

import (
	"fmt"
	"reflect"
)

type (
	Method struct {
	}
)

func (Method) _method() {}

func AssertMethodAnnotation(targetType reflect.Type) {

	switch {
	case !targetType.Implements(reflect.TypeFor[MethodAnnotationConstraint]()):
		panic(
			fmt.Sprintf(
				"Could not apply %s that is not method annotation",
				targetType.Name(),
			),
		)
	}
}
