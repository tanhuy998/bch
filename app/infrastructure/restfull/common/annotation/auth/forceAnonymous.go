package auth

import (
	routeAuth "app/infrastructure/restfull/common/auth/route"
	"app/infrastructure/restfull/common/endpoint/annotation"

	"github.com/kataras/iris/v12/core/router"
)

type (
	ForceAnnonymous struct{ annotation.Annotation }
)

func (ForceAnnonymous) Apply(route *router.Route) {

	routeAuth.Exclude(route)
}
