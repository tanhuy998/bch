package endpoint

import (
	"app/infrastructure/restfull/common/endpoint/internal/session"
	"fmt"
	"reflect"
	"regexp"
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

func __registerEndpoints(reflectTypeCurator reflect.Type, reflectValCurator reflect.Value) {

	methodCount := reflectTypeCurator.NumMethod()

	for i := range methodCount {

		reflectTypeMethod := reflectTypeCurator.Method(i)
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
					`((%s.%s).%s) endpoint method must return 1 value that implements endpoint.IEndpoint, %d value(s) given`,
					reflectTypeCurator.PkgPath(),
					reflectTypeCurator.Name(),
					reflectTypeMethod.Name,
					reflectTypeMethod.Type.NumOut(),
				),
			)
		case !reflectTypeMethod.Type.Out(0).Implements(reflect.TypeFor[IEndpoint]()):
			panic(
				fmt.Sprintf(
					`((%s.%s).%s) endpoint method must return 1 value that implements endpoint.IEndpoint, %s`,
					reflectTypeCurator.PkgPath(),
					reflectTypeCurator.Name(),
					reflectTypeMethod.Name,
					reflectTypeMethod.Type.Out(0).Name(),
				),
			)
		}

		session.Start(
			reflectValCurator, registeredName,
		)
		defer session.End()

		reflectMethod := reflectValCurator.Method(i)

		__prepareAnnotationsAndRetrieveEndpoint(reflectTypeMethod, reflectMethod)
	}
}
