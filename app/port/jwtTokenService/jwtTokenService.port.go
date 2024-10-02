package jwtTokenServicePort

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ERR_SIGNING_METHOD_MISMATCH = errors.New("jwt singing method mismatch")
)

type (
	IJWTTokenGenerator interface {
		GenerateToken() *jwt.Token
	}

	IJWTTokenSigning interface {
		//IJWTTokenGenerator
		SignString(token *jwt.Token) (string, error)
	}

	IJWTTokenVerification interface {
		VerifyTokenStringCustomClaim(token_str string, customClaim jwt.Claims) (*jwt.Token, error)
		VerifyTokenString(string) (*jwt.Token, error)
	}

	IJWTTokenValidator interface {
		Validate(token *jwt.Token) error
	}

	IJWTTokenManipulator interface {
		IJWTTokenValidator
		IJWTTokenGenerator
		IJWTTokenSigning
		IJWTTokenVerification
		GetSigningMethod() interface{}
	}

	ISymmetricJWTTokenManipulator interface {
		IJWTTokenManipulator
		GetSecret() interface{}
	}

	IAsymmetricJWTTokenManipulator interface {
		IJWTTokenManipulator
		GetPrivateKey() interface{}
		GetPublicKey() interface{}
	}
)
