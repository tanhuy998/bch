package jwtTokenService

import (
	jwtTokenServicePort "app/src/port/jwtTokenService"

	"github.com/golang-jwt/jwt/v5"
)

// var (
// 	symmetric_signing_method = jwt.SigningMethodHS256
// )

type (
	jwt_HMACTokenService struct {
		// private_key *ecdsa.PrivateKey
		// public_key  *ecdsa.PublicKey
		secret        []byte
		signingMethod *jwt.SigningMethodHMAC
		//signing_method *jwt.SigningMethodECDSA
	}
)

func NewHMACService(signingMethod *jwt.SigningMethodHMAC, secret []byte) *jwt_HMACTokenService {

	if signingMethod == nil {

		signingMethod = jwt.SigningMethodHS256
	}

	ret := &jwt_HMACTokenService{
		signingMethod: signingMethod,
		secret:        secret,
	}

	return ret
}

func (this *jwt_HMACTokenService) GenerateToken() *jwt.Token {

	return GenerateHMACToken(this.signingMethod)
}

func (this *jwt_HMACTokenService) SignString(token *jwt.Token) (string, error) {

	return token.SignedString(this.secret)
}

func (this *jwt_HMACTokenService) VerifyTokenStringCustomClaim(token_str string, customClaim jwt.Claims) (*jwt.Token, error) {

	return jwt.ParseWithClaims(
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
}

func (this *jwt_HMACTokenService) VerifyTokenString(token_str string) (*jwt.Token, error) {

	return jwt.NewParser(claim_validations...).
		Parse(token_str, func(token *jwt.Token) (interface{}, error) {

			if !IsHMACSigningMethod(token.Method) {

				return nil, jwtTokenServicePort.ERR_SIGNING_METHOD_MISMATCH
			}

			return this.secret, nil
		})
}

func (this *jwt_HMACTokenService) Validate(token *jwt.Token) error {

	return ValidateToken(token)
}

func (this *jwt_HMACTokenService) GetSigningMethod() interface{} {

	return this.signingMethod
}

func (this *jwt_HMACTokenService) GetSecret() interface{} {

	copy := make([]byte, len(this.secret))

	for i, val := range this.secret {

		copy[i] = val
	}

	return copy
}
