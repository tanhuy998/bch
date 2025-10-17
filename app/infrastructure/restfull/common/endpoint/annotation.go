package endpoint

import (
	"app/infrastructure/restfull/common/activator"
	"app/infrastructure/restfull/common/endpoint/annotation"
	"app/infrastructure/restfull/common/endpoint/annotationScope"
	"app/infrastructure/restfull/common/endpoint/internal/annotate"
	"app/infrastructure/restfull/common/endpoint/internal/session"
	"fmt"
	"reflect"

	"github.com/kataras/iris/v12/core/router"
)

type ()

type (

	// ANNOTATION SCOPE
	StructAnnotation interface {
	}

	MethodAnnnotation interface {
	}
	// END ANNOTATION SCOPE

	// ANNOTATION CONSTRAINTS
	SingletonAnnotation interface {
		Singleton()
	}
	PresetSingletonAnnotation interface {
		Singleton() interface{}
	}
	OnceAnnotation interface {
		SingletonAnnotation
		// Once is constraint that annotates the specific annotaiton
		// just affect one time, no matter how it is placed multiple
		// times.
		Once()
	}
	ForceOnceAnnotation interface {
		// The specific annotation have to placed once.
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
func __prepareMethodAnnotationsAndRetrieveEndpoint(
	activator activator.IActivator, reflectTypeMethod reflect.Method, method reflect.Value,
) {

	switch {
	case !session.In():
		panic("endpoint.prepareAndRun() just only been called inside endpoint.RegisterEndpointsOf[T IEndpointBuilder](builder T) function")
	}

	const EXCLUDED_FROM_RECEIVER = 1

	countWithReceiver := reflectTypeMethod.Type.NumIn()
	argCount := countWithReceiver - EXCLUDED_FROM_RECEIVER

	var (
		endpoint    IEndpoint
		annotations []reflect.Value = make([]reflect.Value, argCount)
	)

	defer func() {

		endpoint._buildDone()
	}()

	for _, queuedAccumulator := range annotate.GetAccumulatorQueue() {

		session.Adopt(queuedAccumulator)
	}

	for i := EXCLUDED_FROM_RECEIVER; i < countWithReceiver; i++ {

		type_param := reflectTypeMethod.Type.In(i)

		annotationScope.AssertMethodAnnotation(type_param)

		var (
			challenged           reflect.Value
			shouldSkipInspecting bool
		)

		challenged, shouldSkipInspecting = __prepareAnnotationConstraints(activator, type_param)

		annotations[i-EXCLUDED_FROM_RECEIVER] = challenged

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
	}

	session.ReserveAnnotationsFromEndpoint()
	endpoint = method.Call(annotations)[0].Interface().(IEndpoint)
	endpoint._register()
	session.ReleaseReservation()

	for _, e := range annotate.GetEffectorStack().Snapshot() {

		switch effector := e.Interface().(type) {
		case EndpointEffectorWithAsset:
			effector.Apply(endpoint, session.AssetOf(effector))
		case RouteEffectorWithAsset:
			effector.Apply(endpoint.getRoute(), session.AssetOf(effector))
		case MiddlewareEffectorWithAsset:
			effector.Apply(endpoint, session.AssetOf(effector))
		case EndpointEffector:
			effector.Apply(endpoint)
		case RouteEffector:
			effector.Apply(endpoint.getRoute())
		case MiddlewareEffector:
			effector.Apply(endpoint)
		}
	}
}

func __prepareAnnotationConstraints(
	activator activator.IActivator, annotationType reflect.Type,
) (challenged reflect.Value, shouldSkipInspecting bool) {

	//switch singleton, acknowledgedAsSingleton := annotate.GetSingleton(annotationType); {
	switch singleton, acknowledgedAsSingleton := annotate.GetSingleton(annotationType); {
	case acknowledgedAsSingleton:
		// already registered as singleton
		// validate for two important annotation constraints
		challenged = singleton
		switch challenged.Interface().(type) {
		case ForceOnceAnnotation:
			panic(
				fmt.Sprintf(
					`%s is forced once annotation that ought to be used once`,
					challenged.Type().Name(),
				),
			)
		case OnceAnnotation:
			shouldSkipInspecting = true
		}
		return
	default:
		// not yet acknowledged
		// instantiate new one
		reflectPtr := reflect.New(annotationType)
		// borrow a pointer to wire dependencies of the annotation
		ptr := reflectPtr.Interface()
		activator.Dependencies().Struct(ptr, 0)

		// struct type
		challenged = reflect.ValueOf(ptr).Elem() // reflectPtr.Elem()

		//if type_param.Kind() == reflect.Pointer

		annotation.AssertAnnotation(challenged)

		switch {
		case annotationType.Implements(type_singleton_annotation):
			//singleton_annotations[type_param] = challenged
			annotate.AcknowledgeSingleton(annotationType, challenged)
		case annotationType.Implements(type_preset_singleton_annotation):
			preset := challenged.Interface().(PresetSingletonAnnotation).Singleton()
			//singleton_annotations[type_param] = reflect.ValueOf(stored)
			annotate.AcknowledgeSingleton(annotationType, reflect.ValueOf(preset))
		}

		return
	}
}
