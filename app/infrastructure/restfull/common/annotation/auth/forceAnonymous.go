package auth

import (
	routeAuth "app/infrastructure/restfull/common/auth/route"
	"app/infrastructure/restfull/common/endpoint/annotation"
	"app/infrastructure/restfull/common/endpoint/annotationScope"

	"github.com/kataras/iris/v12/core/router"
)

type (
	ForceAnnonymous struct {
		annotation.Annotation
		annotationScope.Method
		annotationScope.Struct
	}
)

func (ForceAnnonymous) Apply(route *router.Route) {

	routeAuth.Exclude(route)
}
