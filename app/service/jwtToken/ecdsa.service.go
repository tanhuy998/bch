package jwtTokenService

import (
	libError "app/internal/lib/error"
	jwtTokenServicePort "app/port/jwtTokenService"
	"crypto/ecdsa"

	"github.com/golang-jwt/jwt/v5"
)

// var (
// 	signing_method = jwt.SigningMethodES256
// )

type (
	jwt_ECTokenService struct {
		private_key   *ecdsa.PrivateKey
		public_key    *ecdsa.PublicKey
		signingMethod *jwt.SigningMethodECDSA
		//signing_method *jwt.SigningMethodECDSA
	}
)

func NewECDSAService(
	signingMethod *jwt.SigningMethodECDSA,
	privateKey ecdsa.PrivateKey,
	publicKey ecdsa.PublicKey,
) *jwt_ECTokenService {

	if signingMethod == nil {

		signingMethod = jwt.SigningMethodES256
	}

	ret := &jwt_ECTokenService{
		private_key:   &privateKey,
		public_key:    &publicKey,
		signingMethod: signingMethod,
	}

	return ret
}

func (this *jwt_ECTokenService) GenerateToken() *jwt.Token {

	return GenerateECJWTToken(this.signingMethod)
}

func (this *jwt_ECTokenService) SignString(token *jwt.Token) (string, error) {

	ret, err := token.SignedString(this.private_key)

	if err != nil {

		return "", libError.NewInternal(err)
	}

	return ret, nil
}

func (this *jwt_ECTokenService) VerifyTokenStringCustomClaim(token_str string, customClaim jwt.Claims) (*jwt.Token, error) {

	ret, err := jwt.ParseWithClaims(
		token_str,
		customClaim,
		func(token *jwt.Token) (interface{}, error) {

			if !IsECSigningMethod(token.Method) {

				return nil, jwtTokenServicePort.ERR_SIGNING_METHOD_MISMATCH
			}

			return this.public_key, nil
		},
		claim_validations...,
	)

	if err != nil {

		return nil, libError.NewInternal(err)
	}

	return ret, nil
}

func (this *jwt_ECTokenService) VerifyTokenString(token_str string) (*jwt.Token, error) {

	ret, err := jwt.NewParser(claim_validations...).
		Parse(token_str, func(token *jwt.Token) (interface{}, error) {

			if !IsHMACSigningMethod(token.Method) {

				return nil, jwtTokenServicePort.ERR_SIGNING_METHOD_MISMATCH
			}

			return this.public_key, nil
		})

	if err != nil {

		return nil, libError.NewInternal(err)
	}

	return ret, nil
}

func (this *jwt_ECTokenService) Validate(token *jwt.Token) error {

	err := ValidateToken(token)

	if err != nil {

		return libError.NewInternal(err)
	}

	return nil
}

func (this *jwt_ECTokenService) GetSigningMethod() interface{} {

	return this.signingMethod
}

func (this *jwt_ECTokenService) GetPrivateKey() interface{} {

	return *this.private_key
}

func (this *jwt_ECTokenService) GetPublicKey() interface{} {

	return *this.public_key
}
