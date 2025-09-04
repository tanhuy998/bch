package endpoint

import (
	"app/infrastructure/restfull/common/endpoint/internal/session"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
)

const (
	ENDPOINT_PREFIX_CAPTURING_GROUP            = `prefix`
	ENDPOINT_REGISTERED_METHOD_CAPTURING_GROUP = `registerd_method`
	ENDPOINT_PREFIX                            = `ENDPOINT`
)

var (
	regex_match_endpoint_method = regexp.MustCompile(
		fmt.Sprintf(
			`(?P<%s>%s)_(?P<%s>[A-Z][\w_]*)`,
			ENDPOINT_PREFIX_CAPTURING_GROUP,
			ENDPOINT_PREFIX,
			ENDPOINT_REGISTERED_METHOD_CAPTURING_GROUP,
		),
	)
)

func Handle(builder *EndpointBuilder, httpMethod string, path string) IEndpointInitiator {

	switch {
	case !session.In():
		panic("calling endpoint initialization method when not on initialization session is not allowed.")
	case builder == nil:
		panic("bad builder value, nil given")
	case session.RegisteredControllerMethod() == "":
		panic("no registered controller method to handle for endpoint.")
	}

	return NewEnpoint(
		builder.Activator().Handle(
			httpMethod, path, session.RegisteredControllerMethod(),
		),
	)
}

func RegisterEndpointsOf[T IEndpointBuilder](builder T) {

	reflectValBuilder := reflect.ValueOf(any(builder))
	reflectType := reflect.TypeOf(builder)
	methodCount := reflectType.NumMethod()

	for i := 0; i < methodCount; i++ {

		reflectTypeMethod := reflectType.Method(i)
		matches := regex_match_endpoint_method.FindStringSubmatch(reflectTypeMethod.Name)

		switch {
		case len(matches) == 0,
			len(matches) < regex_match_endpoint_method.NumSubexp()+1:
			continue
		}

		prefix := matches[regex_match_endpoint_method.SubexpIndex(ENDPOINT_PREFIX_CAPTURING_GROUP)]

		if prefix != ENDPOINT_PREFIX {

			continue
		}

		registeredName := matches[regex_match_endpoint_method.SubexpIndex(ENDPOINT_REGISTERED_METHOD_CAPTURING_GROUP)]

		if registeredName == "" {

			continue
		}

		switch {
		case reflectTypeMethod.Type.NumOut() != 1:
			panic(
				fmt.Sprintf(
					`(%s.%s) endpoint method must return 1 value that implements endpoint.IEndpoint, %s value(s) given`,
					reflectType.Name(),
					reflectTypeMethod.Name,
					strconv.Itoa(reflectTypeMethod.Type.NumOut()),
				),
			)
		case !reflectTypeMethod.Type.Out(0).Implements(reflect.TypeFor[IEndpoint]()):
			panic(
				fmt.Sprintf(
					`(%s.%s) endpoint method must return 1 value that implements endpoint.IEndpoint, %s`,
					reflectType.Name(),
					reflectTypeMethod.Name,
					reflectTypeMethod.Type.Out(0).Name(),
				),
			)
		}

		session.Start(
			reflectValBuilder, registeredName,
		)

		reflectMethod := reflectValBuilder.Method(i)

		prepareAndCall(reflectTypeMethod, reflectMethod)

		session.End()
	}
}
