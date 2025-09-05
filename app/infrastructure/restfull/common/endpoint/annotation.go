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
)

/*
Read, detect, resolve annations that is placed in the endpoint builder method's structure
*/
func prepareAndCall(reflectTypeMethod reflect.Method, method reflect.Value) {

	switch {
	case !session.In():
		panic("endpoint.prepareAndRun() just only been called inside endpoint.RegisterEndpointsOf[T IEndpointBuilder](builder T) function")
	}

	const passedReceiver = 1

	countWithReceiver := reflectTypeMethod.Type.NumIn()
	argCount := countWithReceiver - passedReceiver

	if countWithReceiver == 1 {
		// first parameter of method reflection is the receiver type reflection, just skip
		method.Call(nil)
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

	for i := passedReceiver; i < countWithReceiver; i++ {

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

				panic(
					fmt.Sprintf(
						`%s is once annotation that ought to be used once`,
						challenged.Type().Name(),
					),
				)
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
