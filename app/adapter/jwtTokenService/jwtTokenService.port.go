package jwtTokenServicePort

import "github.com/golang-jwt/jwt/v5"

type (
	IJWTTokenGenerator interface {
		GenerateToken() *jwt.Token
	}

	IJWTTokenSigning interface {
		//IJWTTokenGenerator
		SignedString(token *jwt.Token) (string, error)
	}

	IJWTTokenVerification interface {
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
	}
)
