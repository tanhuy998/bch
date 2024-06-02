package api

import (
	requestPresenter "app/domain/presenter/request"
	responsePresenter "app/domain/presenter/response"
	"app/internal/controller"
	"app/internal/middleware"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func initCandidateSigningApi(app *iris.Application) *mvc.Application {

	router := app.Party("/sign")

	container := router.ConfigureContainer().Container

	wrapper := mvc.New(router)
	return wrapper.Handle(
		new(controller.SignController),
		applyRoutes(func(activator *mvc.ControllerActivator) {

			/*
				Get Signing info of a candidate
			*/
			activator.Handle(
				"GET", "/campaign/{campaignUUID}/candidate/{candidateUUID}", "GetSingleCandidateSigningInfo",
				middleware.BindPresenters[requestPresenter.GetSingleCandidateSigningInfoRequest, responsePresenter.GetSingleCandidateSigningInfoResponse](container),
			)

			/*
				Post signing info of a candidate
			*/
			activator.Handle(
				"PATCH", "/campaign/{campaignUUID}/candidate/{candidateUUID}", "CommitCandidateSigningInfo",
				middleware.BindPresenters[requestPresenter.CommitCandidateSigningInfoRequest, responsePresenter.CommitCandidateSigningInfoResponse](container),
			)
		}),
	)
}
