package jwtTokenService

import (
	jwtTokenServicePort "app/adapter/jwtTokenService"
)

type (
	IJWTTokenGenerator = jwtTokenServicePort.IJWTTokenGenerator

	IJWTTokenSigning = jwtTokenServicePort.IJWTTokenSigning

	IJWTTokenVerification = jwtTokenServicePort.IJWTTokenVerification

	IJWTTokenValidator = jwtTokenServicePort.IJWTTokenValidator

	IJWTTokenManipulator = jwtTokenServicePort.IJWTTokenManipulator

	IAsymmetricJWTTokenManipulator = jwtTokenServicePort.IAsymmetricJWTTokenManipulator
)
