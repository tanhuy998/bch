package api

import (
	"app/internal/controller"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func initCandidateSigningApi(app *iris.Application) *mvc.Application {

	router := app.Party("/sign")

	wrapper := mvc.New(router)
	return wrapper.Handle(
		new(controller.SignController),
		applyRoutes(func(activator *mvc.ControllerActivator) {

			/*
				Get Signing info of a candidate
			*/
			activator.Handle("GET", "/campaign/{campaignUUID:string}/candidate/{candidateUUID:string}", "GetSigningInfo")

			/*
				Post signing info of a candidate
			*/
			activator.Handle("POST", "/campaign/{campaignUUID:string}/candidate/{candidateUUID:string}", "Sign")
		}),
	)
}
