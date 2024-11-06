package boundedContext

import (
	revokeSignaturesDomain "app/domain/authSignatures/revokeSignatures"
	rotateSignaturesDomain "app/domain/authSignatures/rotateSignatures"
	switchTenantDomain "app/domain/authSignatures/switchTenant"
	libConfig "app/internal/lib/config"
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

	libConfig.BindDependency[accessTokenServicePort.IAccessTokenManipulator, accessTokenService.JWTAccessTokenManipulatorService](container, nil)
	libConfig.BindDependency[accessTokenClientPort.IAccessTokenClient, accessTokenClientService.BearerAccessTokenClientService](container, nil)

	//refreshTokenService := new(refreshTokenService.RefreshTokenManipulatorService)
	libConfig.BindDependency[refreshTokenServicePort.IRefreshTokenManipulator, refreshTokenService.RefreshTokenManipulatorService](container, nil)
	libConfig.BindDependency[refreshTokenClientPort.IRefreshTokenClient, refreshTokenClientService.RefreshTokenClientService](container, nil)

	libConfig.BindDependency[authSignatureTokenPort.IAuthSignatureProvider, authSignatureToken.AuthSignatureTokenService](container, nil)
}

func RegisterAuthSignaturesBoundedContext(container *hero.Container) {

	registerDomainSpecificUtils(container)

	libConfig.BindDependency[authSignaturesServicePort.IRotateSignatures, rotateSignaturesDomain.RotateSignaturesService](container, nil)
	libConfig.BindDependency[authSignaturesServicePort.ISwitchTenant, switchTenantDomain.SwitchTenantService](container, nil)
	libConfig.BindDependency[authSignaturesServicePort.IRevokeSignatures, revokeSignaturesDomain.RevokeSignaturesService](container, nil)

	libConfig.BindDependency[
		usecasePort.IUseCase[requestPresenter.RefreshLoginRequest, responsePresenter.RefreshLoginResponse],
		rotateSignaturesDomain.RotateSignaturesUseCase,
	](container, nil)

	libConfig.BindDependency[
		usecasePort.IUseCase[requestPresenter.SwitchTenant, responsePresenter.SwitchTenant],
		switchTenantDomain.SwitchTenantUseCase,
	](container, nil)

	libConfig.BindDependency[
		usecasePort.IUseCase[requestPresenter.Logout, responsePresenter.Logout],
		revokeSignaturesDomain.RevokeSignaturesUseCase,
	](container, nil)
}
