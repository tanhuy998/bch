package boundedContext

import (
	authenticateCredentialsDomain "app/domain/authGen/authenticateCredentials"
	checkGeneralTokenDomain "app/domain/authGen/checkGeneralToken"
	navigateTenantDomain "app/domain/authGen/navigateTenant"
	irisIoc "app/internal/lib/iris/ioc"

	authGenServicePort "app/port/authGenService"
	usecasePort "app/port/usecase"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"

	"github.com/kataras/iris/v12/hero"
)

func RegisterAuthGenBoundedContext(container *hero.Container) {

	irisIoc.BindDependency[authGenServicePort.IAuthenticateCrdentials, authenticateCredentialsDomain.AuthenticateCredentialsService](container, nil)
	irisIoc.BindDependency[authGenServicePort.INavigateTenant, navigateTenantDomain.NavigateTenantService](container, nil)

	irisIoc.BindDependency[
		usecasePort.IUseCase[requestPresenter.LoginRequest, responsePresenter.LoginResponse],
		authenticateCredentialsDomain.AuthenticateCredentialsUseCase,
	](container, nil)

	irisIoc.BindDependency[
		usecasePort.IUseCase[requestPresenter.CheckLogin, responsePresenter.CheckLogin],
		checkGeneralTokenDomain.CheckGeneralTokenUseCase,
	](container, nil)

	irisIoc.BindDependency[
		usecasePort.IUseCase[requestPresenter.AuthNavigateTenant, responsePresenter.AuthNavigateTenant],
		navigateTenantDomain.NavigateTenantUseCase,
	](container, nil)
}
