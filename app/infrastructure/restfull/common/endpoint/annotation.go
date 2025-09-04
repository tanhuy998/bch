package endpoint

import (
	"app/infrastructure/restfull/common/endpoint/annotation"
	"app/infrastructure/restfull/common/endpoint/internal/session"
	"fmt"
	"reflect"

	"github.com/kataras/iris/v12/core/router"
)

type (
	Accumulator = session.IAccumulator

	// SingletonAccumulator interface {
	// 	Accumulator
	// 	SingletonAnnotation
	// }

	SingletonAnnotation interface {
		Singleton()
	}

	OnceAnnotation interface {
		Once()
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

	RouteEffectorWithAsset interface {
		Accumulator
		Apply(router *router.Route, asset interface{})
	}

	MiddlewareEffectorWithAsset interface {
		Accumulator
		Apply(e IEndpointUseMiddleware, asset interface{})
	}

	EndpointEffectorWithAsset interface {
		Accumulator
		Apply(e IEndpoint, asset interface{})
	}
)

var (
	type_singleton_annotation = reflect.TypeFor[SingletonAnnotation]()

	type_route_effector      = reflect.TypeFor[RouteEffector]()
	type_middleware_effector = reflect.TypeFor[MiddlewareEffector]()
	type_endpoint_effector   = reflect.TypeFor[EndpointEffector]()
)

func assertAffector(t reflect.Type) {

	ok := t.Implements(type_route_effector) ||
		t.Implements(type_middleware_effector) ||
		t.Implements(type_endpoint_effector)

	if !ok {

		panic(
			fmt.Sprintf(`%s is not type of endpoint annotation`, t.Name()),
		)
	}
}

func prepareAnnotationEffector(endpoint *endpoint_builder_t, annotations []reflect.Value) {

	for _, an := range annotations {

		switch eff := an.Interface().(type) {
		case RouteEffector:
			eff.Apply(endpoint.route)
		case MiddlewareEffector:
			eff.Apply(endpoint)
		case EndpointEffector:
			eff.Apply(endpoint)
		}
	}
}

func prepareAndCall(reflectTypeMethod reflect.Method, method reflect.Value) {

	switch {
	case !session.In():
		panic("endpoint.prepareAndRun() just only been called inside endpoint.RegisterEndpointsOf[T IEndpointBuilder](builder T) function")
	}

	const passedReceiver = 1

	count := reflectTypeMethod.Type.NumIn()
	argCount := count - passedReceiver

	if count == 1 {
		// first parameter of method reflection is the receiver type reflection, just skip
		method.Call(make([]reflect.Value, 0))
		return
	}

	annotations := make([]reflect.Value, argCount)

	var (
		singleton_annotations map[reflect.Type]reflect.Value = make(map[reflect.Type]reflect.Value)
		endpoint              IEndpoint
	)

	defer func() {

		singleton_annotations = nil
	}()

	for i := passedReceiver; i < count; i++ {

		type_param := reflectTypeMethod.Type.In(i)

		var challenged reflect.Value

		switch singleton, acknowledgedAsSingleton := singleton_annotations[type_param]; {
		case acknowledgedAsSingleton:
			challenged = singleton
		default:

			challenged = reflect.New(type_param).Elem()

			//if type_param.Kind() == reflect.Pointer

			annotation.AssertAnnotation(challenged)

			if type_param.Implements(type_singleton_annotation) {

				singleton_annotations[type_param] = challenged
			}
		}

		switch challenged.Interface().(type) {
		case OnceAnnotation:
			if _, ok := singleton_annotations[reflectTypeMethod.Type]; ok {

				panic("")
			}
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

		annotations[i-passedReceiver] = challenged
	}

	endpoint = method.Call(annotations)[0].Interface().(IEndpoint) // no need to handle panic
}
