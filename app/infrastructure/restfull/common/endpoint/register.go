package endpoint

import (
	"app/infrastructure/restfull/common/endpoint/internal/annotate"
	"app/infrastructure/restfull/common/endpoint/internal/session"
	"reflect"
)

func Handle(curator *APIEndpointCurator, httpMethod string, path string) IEndpointBuilder {

	switch {
	case !session.In():
		panic("calling endpoint initialization method when not on initialization session is not allowed.")
	case curator == nil:
		panic("bad builder value, nil given")
	case session.RegisteredControllerMethod() == "":
		panic("no registered controller method to handle for endpoint.")
	}

	return NewEnpointBuilder(
		curator.Activator(), httpMethod, path,
	)
}

func RegisterEndpointsOf(curator IAPICurator) {

	LaunchApiOf(curator)
}

func LaunchApiOf(curator IAPICurator) {

	// wire dependencies for controller
	curator.Activator().Dependencies().Struct(curator, 0)

	// curator is pointer
	reflectValCurator := reflect.ValueOf(curator)

	annotate.StackSingletonLayer()
	defer annotate.PopSingletonLayer()

	defer __pop(
		__registerStructAnnotations(curator.Activator(), reflectValCurator),
	)

	defer __pop(
		__registerAnnotations(curator.Activator(), reflectValCurator.Type()),
	)
	__registerMethodEndpoints(curator.Activator(), reflectValCurator)
	__launchRecursiveAPIsOf(curator)
}

func __pop(numAccumulator int, numEffector int) {

	for numAccumulator != 0 || numEffector != 0 {

		switch {
		case numAccumulator > 0:
			annotate.PopAccumulator()
			numAccumulator--
			fallthrough
		case numEffector > 0:
			annotate.PopEffector()
			numEffector--
		}
	}
}

func __launchRecursiveAPIsOf(curator IAPICurator) {

	recursiveCurator, ok := curator.(IRecursiveAPICurator)

	if !ok {

		return
	}

	childAPIs := recursiveCurator.Child()

	if len(childAPIs) == 0 {

		return
	}

	for _, curator := range childAPIs {

		LaunchApiOf(curator)
	}
}
