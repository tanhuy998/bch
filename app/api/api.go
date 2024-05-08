package api

import (
	authService "app/app/service/auth"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

const (
	auth_commander_group = authService.AuthorizationGroup("cmd")
	auth_member_group    = authService.AuthorizationGroup("mbr")
)

func applyRoutes(f func(*mvc.ControllerActivator)) mvc.OptionFunc {

	return f
}

func Init(app *iris.Application) {

	initCandidateSigningApi(app)
	initCampaignGroupApi(app)
	initCandidateGroupApi(app)
}
