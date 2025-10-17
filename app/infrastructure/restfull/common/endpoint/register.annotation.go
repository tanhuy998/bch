package endpoint

import (
	"app/infrastructure/restfull/common/activator"
	"app/infrastructure/restfull/common/endpoint/annotation"
	"app/infrastructure/restfull/common/endpoint/annotationScope"
	"app/infrastructure/restfull/common/endpoint/internal/annotate"
	"reflect"
)

const (
	ANNOTATION_METHOD = "ANNOTATIONS_"
)

func __registerStructAnnotations(activator activator.IActivator, reflectValCurator reflect.Value) (numAccumulator int, numEffector int) {

	if reflectValCurator.Type().Kind() == reflect.Pointer {

		reflectValCurator = reflectValCurator.Elem()
	}

	reflectType := reflectValCurator.Type()

	for i := range reflectValCurator.NumField() {

		reflectTypeField := reflectType.Field(i)
		reflectValField := reflectValCurator.Field(i)

		switch {
		case !reflectTypeField.IsExported(), !annotation.TryAssertAnnotation(reflectValField):
			continue
		default:
			switch isAccumulator, isEffector := reg(activator, reflectTypeField.Type); {
			case isAccumulator:
				numAccumulator++
				fallthrough
			case isEffector:
				numEffector++
			}
		}
	}

	return
}

func __registerAnnotations(activator activator.IActivator, reflectTypeCurator reflect.Type) (numAccumulator int, numEffector int) {

	switch reflectTypeMethod, annotationsMethodExists := reflectTypeCurator.MethodByName(ANNOTATION_METHOD); {
	case !annotationsMethodExists:
		return
	default:
		return __scanAndRegisterAnnotations(activator, reflectTypeMethod)
	}
}

func reg(activator activator.IActivator, reflectTypeAnnotation reflect.Type) (isAccumulator bool, isEffector bool) {

	challenged, shouldSkipInspecting := __prepareAnnotationConstraints(activator, reflectTypeAnnotation)

	if shouldSkipInspecting {
		return
	}

	switch accumulator := challenged.Interface().(type) {
	case Accumulator:
		annotate.StackAccumulator(accumulator)
		isAccumulator = true
	}

	switch challenged.Interface().(type) {
	case EndpointEffectorWithAsset, RouteEffectorWithAsset, MiddlewareEffectorWithAsset,
		EndpointEffector, RouteEffector, MiddlewareEffector:
		annotate.StackEffector(challenged)
		isEffector = true
	}

	return
}

func __scanAndRegisterAnnotations(activator activator.IActivator, reflectTypeMethod reflect.Method) (numAccumulator int, numEffector int) {

	countIncludingReceiver := reflectTypeMethod.Type.NumIn()

	if countIncludingReceiver == 1 {
		// first parameter of method reflection is the receiver type reflection, just skip
		return
	}

	const excludedFromReceiver = 1

	for i := excludedFromReceiver; i < countIncludingReceiver; i++ {

		type_param := reflectTypeMethod.Type.In(i)

		annotationScope.AssertMethodAnnotation(type_param)

		switch isAccumulator, isEffector := reg(activator, type_param); {
		case isAccumulator:
			numAccumulator++
			fallthrough
		case isEffector:
			numEffector++
		}
	}

	return
}
