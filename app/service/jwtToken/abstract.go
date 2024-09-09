package jwtTokenService

import "github.com/golang-jwt/jwt/v5"

type (
	IJWTTokenGenerator interface {
		GenerateToken() *jwt.Token
	}

	IJWTTokenSigning interface {
		SignedString(token *jwt.Token) (string, error)
	}

	IJWTTokenVerification interface {
		VerifyTokenString(string) (*jwt.Token, error)
	}

	IJWTTokenValidator interface {
		Validate(token *jwt.Token) error
	}

	IJWTTokenAuthenticator interface {
		IJWTTokenValidator
		IJWTTokenVerification
	}
)
