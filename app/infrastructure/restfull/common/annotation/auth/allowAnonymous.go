package auth

import (
	"app/infrastructure/restfull/common/endpoint/annotation"
	routeAuth "app/infrastructure/restfull/internal/auth/route"

	"github.com/kataras/iris/v12/core/router"
)

type (
	AllowAnonnymous struct{ annotation.Annotation }
)

func (AllowAnonnymous) Apply(route *router.Route) {

	routeAuth.AllowAnonymous(route)
}
