package auth

import (
	"app/infrastructure/restfull/common/endpoint"
	"app/infrastructure/restfull/common/endpoint/annotation"
	"app/infrastructure/restfull/common/middleware"
	"app/infrastructure/restfull/common/middleware/hook"
	"fmt"
	"time"
)

var (
	authorize_accumulator_key = fmt.Sprintf(`%s %s`, time.Now().String(), "Authorize")
)

type (
	IAnnotationAuthorizeConstraint interface {
		Retrieve() []hook.AuthorityConstraint
	}

	authorize_accumulator_t = []hook.AuthorityConstraint

	Authorize[Constraint IAnnotationAuthorizeConstraint] struct{ annotation.Annotation }
)

func (Authorize[Constraint]) Accumulate(asset interface{}) interface{} {

	if asset == nil {

		return (*new(Constraint)).Retrieve()
	}

	switch constraintChain := asset.(type) {
	case authorize_accumulator_t:
		return append(constraintChain, (*new(Constraint)).Retrieve()...)
	default:
		panic("wrong type of Authorize annotation accumulator, expect []hook.AuthorityConstraint")
	}
}

func (Authorize[Constraint]) GetAccumulatorKey() interface{} {

	return authorize_accumulator_key
}

func (Authorize[Constraint]) Apply(
	endpoint endpoint.IEndpointUseMiddleware, asset interface{},
) {

	switch constraintChain := asset.(type) {
	case []hook.AuthorityConstraint:
		endpoint.Middleware(
			middleware.Auth(
				constraintChain...,
			),
		)
	}
}
