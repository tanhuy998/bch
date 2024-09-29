package jwtTokenService

import (
	jwtTokenServicePort "app/adapter/jwtTokenService"

	"github.com/golang-jwt/jwt/v5"
)

var (
	claim_validations []jwt.ParserOption = []jwt.ParserOption{
		jwt.WithoutClaimsValidation(),
	}
)

type (
	IJWTTokenGenerator = jwtTokenServicePort.IJWTTokenGenerator

	IJWTTokenSigning = jwtTokenServicePort.IJWTTokenSigning

	IJWTTokenVerification = jwtTokenServicePort.IJWTTokenVerification

	IJWTTokenValidator = jwtTokenServicePort.IJWTTokenValidator

	IJWTTokenManipulator = jwtTokenServicePort.IJWTTokenManipulator

	IAsymmetricJWTTokenManipulator = jwtTokenServicePort.IAsymmetricJWTTokenManipulator
)
