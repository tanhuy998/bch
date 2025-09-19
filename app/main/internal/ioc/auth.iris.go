package ioc

import (
	checkAuthorityDomain "app/domain/auth/checkAuthority"
	irisIoc "app/internal/lib/iris/ioc"
	"app/main/internal/dependencies/jwt"
	"app/main/internal/dependencies/log"

	accessTokenServicePort "app/port/accessToken"
	accessTokenClientPort "app/port/accessTokenClient"
	"app/port/authenticatorServicePort"
	generalTokenServicePort "app/port/generalToken"
	generalTokenClientServicePort "app/port/generalTokenClient"
	generalTokenIDServicePort "app/port/generalTokenID"
	usecasePort "app/port/usecase"

	jwtTokenServicePort "app/port/jwtTokenService"
	refreshTokenIdServicePort "app/port/refreshTokenID"
	uniqueIDServicePort "app/port/uniqueID"
	accessTokenClientService "app/service/accessTokenClient"
	"app/service/accessTokenService"
	"app/service/authenticatorService"
	generalTokenClientService "app/service/generalTokenClient"
	generalTokenIDService "app/service/generalTokenID"
	"app/service/generalTokenService"
	refreshTokenIDService "app/service/refreshTokenID"
	uniqueIDService "app/service/uniqueID"

	"github.com/kataras/iris/v12/hero"
)

func RegisterAuthDependencies(container *hero.Container) {

	log.Main().Println("Initialize Auth service...")

	//irisIoc.BindAndMapDependencyToContext[authService.IAuthService, authService.AuthenticationService](container, nil, AUTH)

	// asymmetricJWTService := jwtTokenService.NewECDSAService(
	// 	jwt.SigningMethodES256, *bootstrap.GetJWTAsymmetricEncryptionPrivateKey(), *bootstrap.GetJWTAsymmetricEncryptionPublicKey(),
	// )
	// irisIoc.BindDependency[jwtTokenServicePort.IAsymmetricJWTTokenManipulator](container, asymmetricJWTService)
	irisIoc.BindDependency[jwtTokenServicePort.IAsymmetricJWTTokenManipulator](container, jwt.NewAsymetricJWTService())

	// symmetricJWTService := jwtTokenService.NewHMACService(
	// 	jwt.SigningMethodHS256, bootstrap.GetJWTSymmetricEncryptionSecret(),
	// )
	// irisIoc.BindDependency[jwtTokenServicePort.ISymmetricJWTTokenManipulator](container, symmetricJWTService)
	irisIoc.BindDependency[jwtTokenServicePort.ISymmetricJWTTokenManipulator](container, jwt.NewSymetricJWTService())

	uniqueID, err := uniqueIDService.New(15)

	if err != nil {

		panic("error while initiating uniqueID service: " + err.Error())
	}

	irisIoc.BindDependency[accessTokenServicePort.IAccessTokenReader, accessTokenService.AccessTokenManufacturerService](container, nil)

	irisIoc.BindDependency[uniqueIDServicePort.IUniqueIDGenerator](container, uniqueID)
	irisIoc.BindDependency[generalTokenIDServicePort.IGeneralTokenIDProvider, generalTokenIDService.GeneralTokenIDProvider](container, nil)

	irisIoc.BindDependency[refreshTokenIdServicePort.IRefreshTokenIDProvider, refreshTokenIDService.RefreshTokenIDProviderService](container, nil)

	irisIoc.BindDependency[generalTokenServicePort.IGeneralTokenManipulator, generalTokenService.GeneralTokenManipulator](container, nil)
	irisIoc.BindDependency[generalTokenClientServicePort.IGeneralTokenClient, generalTokenClientService.GeneralTokenClientService](container, nil)
	irisIoc.BindDependency[accessTokenClientPort.IAccessTokenClient, accessTokenClientService.BearerAccessTokenClientService](container, nil)

	irisIoc.BindDependency[
		usecasePort.IMiddlewareUseCase, checkAuthorityDomain.CheckAuthoritySessionUseCase,
	](container, nil)

	irisIoc.BindDependency[authenticatorServicePort.IAuthenticator, authenticatorService.AuthenticatorService](container, nil)
	log.Main().Println("Auth service initialized.")
}
