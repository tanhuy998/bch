package authGenApi

import (
	"app/infrastructure/http/common"
	"app/infrastructure/restfull/common/annotation/auth"
	"app/infrastructure/restfull/common/annotation/input"
	"app/infrastructure/restfull/common/crud"
	"app/infrastructure/restfull/common/endpoint"

	"app/model"
	usecasePort "app/port/usecase"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"

	"github.com/kataras/iris/v12/mvc"
)

type (
	AuthGeneralController struct {
		crud.EndpointCurator
		common.Controller
		AuthenticateCredentialsUseCase usecasePort.IUseCase[requestPresenter.LoginRequest, responsePresenter.LoginResponse]
		NavigateTenantUseCase          usecasePort.IUseCase[requestPresenter.AuthNavigateTenant, responsePresenter.AuthNavigateTenant[model.Tenant]]
		CheckGenTokenUseCase           usecasePort.IUseCase[requestPresenter.CheckLogin, responsePresenter.CheckLogin]
	}
)

func (this *AuthGeneralController) ANNOTATIONS_(
	auth.ForceAnnonymous,
) {
}

func (this *AuthGeneralController) ENDPOINT_AuthenticateCredentials(
	input.Bind[requestPresenter.LoginRequest],
) endpoint.IEndpoint {
	return this.POST("/credentials").
		BuildAction(
			func(input *requestPresenter.LoginRequest) (mvc.Result, error) {

				return this.ResultOf(
					this.AuthenticateCredentialsUseCase.Execute(input),
				)
			},
		)
}

func (this *AuthGeneralController) ENDPOINT_NavigateTenant(
	input.Bind[requestPresenter.AuthNavigateTenant],
) endpoint.IEndpoint {
	return this.GET("/nav").
		BuildAction(
			func(input *requestPresenter.AuthNavigateTenant) (mvc.Result, error) {

				return this.ResultOf(
					this.NavigateTenantUseCase.Execute(input),
				)
			},
		)
}

func (this *AuthGeneralController) ENDPOINT_CheckGenToken(
	input.Bind[requestPresenter.CheckLogin],
) endpoint.IEndpoint {
	return this.HEAD("/").
		BuildAction(
			func(input *requestPresenter.CheckLogin) (mvc.Result, error) {

				return this.ResultOf(
					this.CheckGenTokenUseCase.Execute(input),
				)
			},
		)
}
