package jwtTokenService

import (
	libError "app/internal/lib/error"
	jwtTokenServicePort "app/port/jwtTokenService"

	"github.com/golang-jwt/jwt/v5"
)

// var (
// 	symmetric_signing_method = jwt.SigningMethodHS256
// )

type (
	JWTHMACTokenService struct {
		// private_key *ecdsa.PrivateKey
		// public_key  *ecdsa.PublicKey
		secret        []byte
		signingMethod *jwt.SigningMethodHMAC
		//signing_method *jwt.SigningMethodECDSA
	}
)

func NewHMACService(signingMethod *jwt.SigningMethodHMAC, secret []byte) *JWTHMACTokenService {

	if signingMethod == nil {

		signingMethod = jwt.SigningMethodHS256
	}

	ret := &JWTHMACTokenService{
		signingMethod: signingMethod,
		secret:        secret,
	}

	return ret
}

func (this *JWTHMACTokenService) GenerateToken() *jwt.Token {

	return GenerateHMACToken(this.signingMethod)
}

func (this *JWTHMACTokenService) SignString(token *jwt.Token) (string, error) {

	ret, err := token.SignedString(this.secret)

	if err != nil {

		return "", libError.NewInternal(err)
	}

	return ret, nil
}

func (this *JWTHMACTokenService) VerifyTokenStringCustomClaim(token_str string, customClaim jwt.Claims) (*jwt.Token, error) {

	ret, err := jwt.ParseWithClaims(
		token_str,
		customClaim,
		func(token *jwt.Token) (interface{}, error) {

			if !IsHMACSigningMethod(token.Method) {

				return nil, jwtTokenServicePort.ERR_SIGNING_METHOD_MISMATCH
			}

			return this.secret, nil
		},
		claim_validations...,
	)

	if err != nil {

		return nil, libError.NewInternal(err)
	}

	return ret, nil
}

func (this *JWTHMACTokenService) VerifyTokenString(token_str string) (*jwt.Token, error) {

	ret, err := jwt.NewParser(claim_validations...).
		Parse(token_str, func(token *jwt.Token) (interface{}, error) {

			if !IsHMACSigningMethod(token.Method) {

				return nil, jwtTokenServicePort.ERR_SIGNING_METHOD_MISMATCH
			}

			return this.secret, nil
		})

	if err != nil {

		return nil, libError.NewInternal(err)
	}

	return ret, nil
}

func (this *JWTHMACTokenService) Validate(token *jwt.Token) error {

	err := ValidateToken(token)

	if err != nil {

		return libError.NewInternal(err)
	}

	return nil
}

func (this *JWTHMACTokenService) GetSigningMethod() interface{} {

	return this.signingMethod
}

func (this *JWTHMACTokenService) GetSecret() interface{} {

	copy := make([]byte, len(this.secret))

	for i, val := range this.secret {

		copy[i] = val
	}

	return copy
}
