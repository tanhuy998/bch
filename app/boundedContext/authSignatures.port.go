package boundedContext

import (
	revokeSignaturesDomain "app/domain/authSignatures/revokeSignatures"
	rotateSignaturesDomain "app/domain/authSignatures/rotateSignatures"
	switchTenantDomain "app/domain/authSignatures/switchTenant"
	irisIoc "app/internal/lib/iris/ioc"

	accessTokenServicePort "app/port/accessToken"
	accessTokenClientPort "app/port/accessTokenClient"
	authSignatureTokenPort "app/port/authSignatureToken"
	authSignaturesServicePort "app/port/authSignatures"
	refreshTokenServicePort "app/port/refreshToken"
	refreshTokenClientPort "app/port/refreshTokenClient"
	usecasePort "app/port/usecase"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"
	accessTokenClientService "app/service/accessTokenClient"
	"app/service/accessTokenService"
	"app/service/authSignatureToken"
	refreshTokenService "app/service/refreshToken"
	refreshTokenClientService "app/service/refreshTokenClient"

	"github.com/kataras/iris/v12/hero"
)

func registerDomainSpecificUtils(container *hero.Container) {

	irisIoc.BindDependency[accessTokenServicePort.IAccessTokenManipulator, accessTokenService.JWTAccessTokenManipulatorService](container, nil)
	irisIoc.BindDependency[accessTokenClientPort.IAccessTokenClient, accessTokenClientService.BearerAccessTokenClientService](container, nil)

	//refreshTokenService := new(refreshTokenService.RefreshTokenManipulatorService)
	irisIoc.BindDependency[refreshTokenServicePort.IRefreshTokenManipulator, refreshTokenService.RefreshTokenManipulatorService](container, nil)
	irisIoc.BindDependency[refreshTokenClientPort.IRefreshTokenClient, refreshTokenClientService.RefreshTokenClientService](container, nil)

	irisIoc.BindDependency[authSignatureTokenPort.IAuthSignatureProvider, authSignatureToken.AuthSignatureTokenService](container, nil)
}

func RegisterAuthSignaturesBoundedContext(container *hero.Container) {

	registerDomainSpecificUtils(container)

	irisIoc.BindDependency[authSignaturesServicePort.IRotateSignatures, rotateSignaturesDomain.RotateSignaturesService](container, nil)
	irisIoc.BindDependency[authSignaturesServicePort.ISwitchTenant, switchTenantDomain.SwitchTenantService](container, nil)
	irisIoc.BindDependency[authSignaturesServicePort.IRevokeSignatures, revokeSignaturesDomain.RevokeSignaturesService](container, nil)

	irisIoc.BindDependency[
		usecasePort.IUseCase[requestPresenter.RefreshLoginRequest, responsePresenter.RefreshLoginResponse],
		rotateSignaturesDomain.RotateSignaturesUseCase,
	](container, nil)

	irisIoc.BindDependency[
		usecasePort.IUseCase[requestPresenter.SwitchTenant, responsePresenter.SwitchTenant],
		switchTenantDomain.SwitchTenantUseCase,
	](container, nil)

	irisIoc.BindDependency[
		usecasePort.IUseCase[requestPresenter.Logout, responsePresenter.Logout],
		revokeSignaturesDomain.RevokeSignaturesUseCase,
	](container, nil)
}
