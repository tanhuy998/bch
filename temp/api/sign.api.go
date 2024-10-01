package api

import (
	"app/src/infrastructure/http/api/v1/controller"

	"github.com/kataras/iris/v12/core/router"
	"github.com/kataras/iris/v12/mvc"
)

func initCandidateSigningApi(app router.Party) *mvc.Application {

	router := app.Party("/signing")

	//container := router.ConfigureContainer().Container

	wrapper := mvc.New(router)
	return wrapper.Handle(
		new(controller.SignController),
		// applyRoutes(func(activator *mvc.ControllerActivator) {

		// 	/*
		// 		Get Signing info of a candidate
		// 	*/
		// 	activator.Handle(
		// 		"GET", "/candidate/{candidateUUID}", "GetSingleCandidateSigningInfo",
		// 		middleware.BindPresenters[requestPresenter.GetSingleCandidateSigningInfoRequest, responsePresenter.GetSingleCandidateSigningInfoResponse](container),
		// 	)

		// 	/*
		// 		Post signing info of a candidate
		// 	*/
		// 	activator.Handle(
		// 		"PATCH", "/campaign/{campaignUUID}/candidate/{candidateUUID}", "CommitCandidateSigningInfo",
		// 		middleware.BindPresenters[requestPresenter.CommitCandidateSigningInfoRequest, responsePresenter.CommitCandidateSigningInfoResponse](container),
		// 	)

		// 	activator.Handle(
		// 		"HEAD", "/pending/campaign/{campaignUUID}/candidate/{candidateUUID}", "CheckSigningExistence",
		// 		middleware.BindPresenters[requestPresenter.CheckSigningExistenceRequest, presenter.IEmptyPresenter](container),
		// 	)

		// 	activator.Handle(
		// 		"PUT", "/campaign/{campaignUUID}/candidate/{candidateUUID}", "CommitSpecificSigningInfo",
		// 		middleware.BindPresenters[requestPresenter.CommitSpecificSigningInfo, responsePresenter.CommitSpecificSigningInfoResponse](container),
		// 	)

		// }),
	)
}
