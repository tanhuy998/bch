package auth

import (
	"app/infrastructure/restfull/common/endpoint/annotation"
	routeAuth "app/infrastructure/restfull/internal/auth/route"

	"github.com/kataras/iris/v12/core/router"
)

type (
	NoAuthorize struct{ annotation.Annotation }
)

func (NoAuthorize) Apply(route *router.Route) {

	routeAuth.NoAuthorize(route)
}
