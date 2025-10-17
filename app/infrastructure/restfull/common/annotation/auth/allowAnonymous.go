package auth

import (
	"app/infrastructure/restfull/common/endpoint/annotation"
	"app/infrastructure/restfull/common/endpoint/annotationScope"
	routeAuth "app/infrastructure/restfull/internal/auth/route"

	"github.com/kataras/iris/v12/core/router"
)

type (
	AllowAnonnymous struct {
		annotation.Annotation
		annotationScope.Method
		annotationScope.Struct
	}
)

func (AllowAnonnymous) Apply(route *router.Route) {

	routeAuth.AllowAnonymous(route)
}
