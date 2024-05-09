package api

import (
	"app/app/controller"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func initCandidateSigningApi(app *iris.Application) {

	router := app.Party("/sign")

	wrapper := mvc.New(router)
	wrapper.Handle(
		new(controller.SignController),
		applyRoutes(func(activator *mvc.ControllerActivator) {

			activator.Handle("GET", "/campaign/{campaignUUID:string}/candidate/{candidateUUID:string}", "GetSigningInfo")
			activator.Handle("POST", "/", "Sign")
		}),
	)
}
