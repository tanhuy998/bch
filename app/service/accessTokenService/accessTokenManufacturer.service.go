package accessTokenService

import (
	libError "app/internal/lib/error"
	jwtTokenServicePort "app/port/jwtTokenService"
)

type (
	AccessTokenManufacturerService struct {
		JWTTokenManipulatorService jwtTokenServicePort.IAsymmetricJWTTokenManipulator
	}
)

func (this *AccessTokenManufacturerService) Read(token_str string) (IAccessToken, error) {

	token, err := this.JWTTokenManipulatorService.VerifyTokenStringCustomClaim(token_str, &jwt_access_token_custom_claims{})

	if err != nil {

		return nil, err
	}

	return newFromToken(token)
}

func (this *AccessTokenManufacturerService) SignString(accessToken IAccessToken) (string, error) {

	if val, ok := accessToken.(*jwt_access_token); ok {

		return this.JWTTokenManipulatorService.SignString(val.jwt_token)
	}

	return "", libError.NewInternal(ERR_INVALID_ACCESS_TOKEN_TYPE)
}
