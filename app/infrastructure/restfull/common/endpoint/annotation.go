package endpoint

import (
	"app/infrastructure/restfull/common/endpoint/annotation"
	"app/infrastructure/restfull/common/endpoint/internal/annotate"
	"app/infrastructure/restfull/common/endpoint/internal/session"
	"fmt"
	"reflect"

	"github.com/kataras/iris/v12/core/router"
)

type (
	// ANNOTATION CONSTRAINTS
	SingletonAnnotation interface {
		Singleton()
	}
	PresetSingletonAnnotation interface {
		Singleton() interface{}
	}
	OnceAnnotation interface {
		Once()
	}
	ForceOnceAnnotation interface {
		OnceAnnotation
		Panic()
	}
	// END ANNOTATION CONSTRAINTS

	// PRE-ENDPOINT ANNOTATION
	Accumulator = session.IAccumulator
	// END PRE-ENDPOINT ANNOTATIONS

	// POST-END POINT ANNOTATIONS
	EndpointEffectorWithAsset interface {
		Accumulator
		Apply(e IEndpoint, asset interface{})
	}
	RouteEffectorWithAsset interface {
		Accumulator
		Apply(router *router.Route, asset interface{})
	}
	MiddlewareEffectorWithAsset interface {
		Accumulator
		Apply(e IEndpointUseMiddleware, asset interface{})
	}
	RouteEffector interface {
		Apply(*router.Route)
	}
	MiddlewareEffector interface {
		Apply(IEndpointUseMiddleware)
	}
	EndpointEffector interface {
		Apply(IEndpoint)
	}
	// END POST-ENDPOINT ANNOTATIONS
)

var (
	type_singleton_annotation        = reflect.TypeFor[SingletonAnnotation]()
	type_preset_singleton_annotation = reflect.TypeFor[PresetSingletonAnnotation]()
)

/*
Read, detect, resolve annations that is placed in the endpoint builder method's structure
*/
func __prepareAnnotationsAndRetrieveEndpoint(
	reflectTypeMethod reflect.Method, method reflect.Value,
) {

	switch {
	case !session.In():
		panic("endpoint.prepareAndRun() just only been called inside endpoint.RegisterEndpointsOf[T IEndpointBuilder](builder T) function")
	}

	const excludedFromReceiver = 1

	countWithReceiver := reflectTypeMethod.Type.NumIn()
	argCount := countWithReceiver - excludedFromReceiver

	var (
		endpoint    IEndpoint
		annotations []reflect.Value = make([]reflect.Value, argCount)
	)

	for _, queuedAccumulator := range annotate.GetAccumulatorQueue() {

		session.Adopt(queuedAccumulator)
	}

	for i := excludedFromReceiver; i < countWithReceiver; i++ {

		type_param := reflectTypeMethod.Type.In(i)

		var (
			challenged           reflect.Value
			shouldSkipInspecting bool
		)

		challenged, shouldSkipInspecting = __prepareAnnotationConstraints(type_param)

		if shouldSkipInspecting {

			continue
		}

		switch accumulator := challenged.Interface().(type) {
		case Accumulator:
			session.Adopt(accumulator)
		}

		switch effector := challenged.Interface().(type) {
		case EndpointEffectorWithAsset:
			defer func() {
				effector.Apply(endpoint, session.AssetOf(effector))
			}()
		case RouteEffectorWithAsset:
			defer func() {
				effector.Apply(endpoint.getRoute(), session.AssetOf(effector))
			}()
		case MiddlewareEffectorWithAsset:
			defer func() {
				effector.Apply(endpoint, session.AssetOf(effector))
			}()
		case EndpointEffector:
			defer func() {
				effector.Apply(endpoint)
			}()
		case RouteEffector:
			defer func() {
				effector.Apply(endpoint.getRoute())
			}()
		case MiddlewareEffector:
			defer func() {
				effector.Apply(endpoint)
			}()
		}

		annotations[i-excludedFromReceiver] = challenged
	}

	session.ReserveAnnotationsFromEndpoint()
	defer session.ReleaseReservation()
	endpoint = method.Call(annotations)[0].Interface().(IEndpoint)
	endpoint._register()
}

func __prepareAnnotationConstraints(type_param reflect.Type) (challenged reflect.Value, shouldSkipInspecting bool) {

	switch singleton, acknowledgedAsSingleton := annotate.GetSingleton(type_param); {
	case acknowledgedAsSingleton:
		challenged = singleton
	default:

		challenged = reflect.New(type_param).Elem()

		//if type_param.Kind() == reflect.Pointer

		annotation.AssertAnnotation(challenged)

		switch {
		case type_param.Implements(type_singleton_annotation):
			//singleton_annotations[type_param] = challenged
			annotate.AcknowledgeSingleton(type_param, challenged)
		case type_param.Implements(type_preset_singleton_annotation):
			preset := challenged.Interface().(PresetSingletonAnnotation).Singleton()
			//singleton_annotations[type_param] = reflect.ValueOf(stored)
			annotate.AcknowledgeSingleton(type_param, reflect.ValueOf(preset))
		}
	}

	switch challenged.Interface().(type) {
	case ForceOnceAnnotation:
		if annotate.HasSingleton(type_param) {
			panic(
				fmt.Sprintf(
					`%s is forced once annotation that ought to be used once`,
					challenged.Type().Name(),
				),
			)
		}
		annotate.AcknowledgeSingleton(type_param, challenged)
	case OnceAnnotation:
		if annotate.HasSingleton(type_param) {
			shouldSkipInspecting = true
		} else {
			annotate.AcknowledgeSingleton(type_param, challenged)
		}
	}

	return
}
