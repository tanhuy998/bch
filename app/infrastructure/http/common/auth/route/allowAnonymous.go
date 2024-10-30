/*
	Define auth excluded feature for routes
	do not use behind parent route that use middleware.Auth
	with authority constraints.
*/

package routeAuth

import (
	commonAuth "app/infrastructure/http/common/auth/internal"

	"github.com/kataras/iris/v12/core/router"
)

/*
*
AllowAnonymous means both authenticated request with access token
and request without access token could be passed the middleware.Auth
*/
func AllowAnonymous(route *router.Route) {

	commonAuth.ExcludePath(
		route.Method + route.Tmpl().Src,
	)
}

/*
Exclude meams the request must contains no access token
if access token comes along, this request is known as bad request
*/
func Exclude(route *router.Route) {

	commonAuth.MarkAnonymous(
		route.Method + route.Tmpl().Src,
	)
}
