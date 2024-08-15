package api

import (
	authService "app/service/auth"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

const (
	auth_commander_group = authService.AuthorizationGroup("cmd")
	auth_member_group    = authService.AuthorizationGroup("mbr")
	auth_delete_claim    = authService.AuthorizationClaim("remove_claim")
	auth_post_claim      = authService.AuthorizationClaim("post_claim")
	auth_get_claim       = authService.AuthorizationClaim("get_claim")
	auth_put_claim       = authService.AuthorizationClaim("put_claim")
)

func applyRoutes(f func(*mvc.ControllerActivator)) mvc.OptionFunc {

	return f
}

func Init(app *iris.Application) {

	initCandidateSigningApi(app).EnableStructDependents()
	initCampaignGroupApi(app).EnableStructDependents()
	initCandidateGroupApi(app).EnableStructDependents()
	initAuthApi(app)
}
