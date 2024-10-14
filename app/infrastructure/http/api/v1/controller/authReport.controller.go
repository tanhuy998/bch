package controller

import (
	"app/infrastructure/http/common"
	"app/infrastructure/http/middleware"
	"app/infrastructure/http/middleware/middlewareHelper"
	usecasePort "app/port/usecase"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"

	"github.com/kataras/iris/v12/mvc"
)

type (
	AuthReportController struct {
		common.Controller
		ReportParitcipatedCommandGroupsUseCase usecasePort.IUseCase[requestPresenter.ReportParticipatedGroups, responsePresenter.ReportParticipatedGroups]
	}
)

func (this *AuthReportController) BeforeActivation(activator mvc.BeforeActivation) {

	container := activator.Dependencies()

	activator.Router().Use(
		middleware.Auth(
			container,
		),
	)

	activator.Handle(
		"GET", "/groups/participated/user/{userUUID:uuid}", "ReportParticipatedGroups",
		middleware.Auth(
			container,
			middlewareHelper.AuthRequiredTenantAgentExceptMeetRoles("COMMANDER"),
		),
		middleware.BindRequest[requestPresenter.ReportParticipatedGroups](
			container,
			middlewareHelper.UseAuthority,
			middlewareHelper.UseTenantMapping,
		),
	)
}

func (this *AuthReportController) ReportParticipatedGroups(
	input *requestPresenter.ReportParticipatedGroups,
) (mvc.Result, error) {

	return this.ResultOf(
		this.ReportParitcipatedCommandGroupsUseCase.Execute(input),
	)
}
