package controller

import (
	"app/infrastructure/restfull/common"
	"app/infrastructure/restfull/common/annotation/input"
	"app/infrastructure/restfull/common/crud"
	"app/infrastructure/restfull/common/endpoint"
	usecasePort "app/port/usecase"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"

	"github.com/kataras/iris/v12/mvc"
)

type (
	SignatruresController struct {
		common.Controller
		crud.EndpointCurator
		RetrieveTenantSignaturesUseCase usecasePort.IUseCase[requestPresenter.SwitchTenant, responsePresenter.SwitchTenant]
		RefreshTenantSignaturesUseCase  usecasePort.IUseCase[requestPresenter.RefreshLoginRequest, responsePresenter.RefreshLoginResponse]
		RevokeTenantSignaturesUseCase   usecasePort.IUseCase[requestPresenter.Logout, responsePresenter.Logout]
	}
)

func (this *SignatruresController) ENDPOINT_RetrieveTenantSignatures(
	input.Bind[requestPresenter.SwitchTenant],
) endpoint.IEndpoint {
	return this.GET("/tenant/{tenantUUID:uuid}").
		BuildAction(
			func(input *requestPresenter.SwitchTenant) (mvc.Result, error) {

				return this.ResultOf(
					this.RetrieveTenantSignaturesUseCase.Execute(input),
				)
			},
		)
}

func (this *SignatruresController) ENDPOINT_RefreshTenantSignatures(
	input.Bind[requestPresenter.RefreshLoginRequest],
) endpoint.IEndpoint {
	return this.POST("/").
		BuildAction(
			func(input *requestPresenter.RefreshLoginRequest) (mvc.Result, error) {

				return this.ResultOf(
					this.RefreshTenantSignaturesUseCase.Execute(input),
				)
			},
		)
}

func (this *SignatruresController) ENDPOINT_RevokeTenantSignatures(
	input.Bind[requestPresenter.Logout],
) endpoint.IEndpoint {
	return this.DELETE("/").
		BuildAction(
			func(input *requestPresenter.Logout) (mvc.Result, error) {

				return this.ResultOf(
					this.RevokeTenantSignaturesUseCase.Execute(input),
				)
			},
		)
}
