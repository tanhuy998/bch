package endpoint

import (
	"app/infrastructure/restfull/common/endpoint/internal/annotate"
	"reflect"
)

const (
	ANNOTATION_METHOD = "ANNOTATIONS_"
)

func __registerAnnotations(reflectTypeCurator reflect.Type) (numAccumulator int, numEffector int) {

	switch reflectTypeMethod, annotationsMethodExists := reflectTypeCurator.MethodByName(ANNOTATION_METHOD); {
	case !annotationsMethodExists:
		return
	default:
		return __scanAndRegisterAnnotations(reflectTypeMethod)
	}

}

func __scanAndRegisterAnnotations(reflectTypeMethod reflect.Method) (numAccumulator int, numEffector int) {

	countIncludingReceiver := reflectTypeMethod.Type.NumIn()

	if countIncludingReceiver == 1 {
		// first parameter of method reflection is the receiver type reflection, just skip
		return
	}

	const excludedFromReceiver = 1

	for i := excludedFromReceiver; i < countIncludingReceiver; i++ {

		type_param := reflectTypeMethod.Type.In(i)

		var challenged reflect.Value

		challenged, _ = __prepareAnnotationConstraints(type_param)

		switch accumulator := challenged.Interface().(type) {
		case Accumulator:
			annotate.StackAccumulator(accumulator)
		}

		switch challenged.Interface().(type) {
		case EndpointEffectorWithAsset, RouteEffectorWithAsset, MiddlewareEffectorWithAsset,
			EndpointEffector, RouteEffector, MiddlewareEffector:
			annotate.StackEffector(challenged)
		}
	}

	return
}
