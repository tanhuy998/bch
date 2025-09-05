package auth

import (
	"app/infrastructure/restfull/common/endpoint/annotation"
	"app/infrastructure/restfull/common/middleware/hook"
	routeAuth "app/infrastructure/restfull/internal/auth/route"
	"fmt"
	"time"

	"github.com/kataras/iris/v12/core/router"
)

type (
	AllowAnonnymous struct{ annotation.Annotation }
)

func (AllowAnonnymous) Apply(route *router.Route) {

	routeAuth.AllowAnonymous(route)
}

type (
	Annonymous struct{ annotation.Annotation }
)

func (Annonymous) Apply(route *router.Route) {

	routeAuth.Exclude(route)
}

type (
	NoAuthorize struct{ annotation.Annotation }
)

func (NoAuthorize) Apply(route *router.Route) {

	routeAuth.NoAuthorize(route)
}

var (
	authorize_accumulator_key = fmt.Sprintf(`%s %s`, time.Now().String(), "Authorize")
)

type (
	IAuthorityConstraint interface {
		Retrieve() []hook.AuthorityConstraint
	}

	authorize_accumulator_t = []hook.AuthorityConstraint

	Authorize[Constraint IAuthorityConstraint] struct{ annotation.Annotation }
)

func (Authorize[Constraint]) Accumulate(asset interface{}) interface{} {

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
