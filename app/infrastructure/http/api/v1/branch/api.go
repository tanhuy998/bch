package api

import (
	"app/infrastructure/http/middleware"

	"github.com/kataras/iris/v12"
)

// const (
// 	auth_commander_group = authService.AuthorizationGroup("cmd")
// 	auth_member_group    = authService.AuthorizationGroup("mbr")
// 	auth_delete_claim    = authService.AuthorizationClaim("remove_claim")
// 	auth_post_claim      = authService.AuthorizationClaim("post_claim")
// 	auth_get_claim       = authService.AuthorizationClaim("get_claim")
// 	auth_put_claim       = authService.AuthorizationClaim("put_claim")
// )

// func applyRoutes(f func(*mvc.ControllerActivator)) mvc.OptionFunc {

// 	return f
// }

func Init(app *iris.Application) {

	container := app.ConfigureContainer().Container

	app.UseRouter(middleware.InternalAccessLog(container))

	initTenantApi(app).EnableStructDependents()

	//tenantIsolationRouter := app.Party("/tenant/{tenantUUID:uuid}")

	//initCandidateSigningApi(app).EnableStructDependents()
	//initCampaignGroupApi(app).EnableStructDependents()
	//initCandidateGroupApi(app).EnableStructDependents()
	initAssignmentApi(app).EnableStructDependents()
	//initInternalAPI(app)
	initAuthApi(app)
}
