package jwtTokenService

import (
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

	return token.SignedString(this.private_key)
}

func (this *jwt_ECTokenService) VerifyTokenString(token_str string) (*jwt.Token, error) {

	return jwt.Parse(token_str, func(token *jwt.Token) (interface{}, error) {

		// if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {

		// 	return nil, ERR_SIGNING_METHOD_MISMATCH
		// }

		if !IsECSigningMethod(token.Method) {

			return nil, ERR_SIGNING_METHOD_MISMATCH
		}

		return this.public_key, nil
	})
}

func (this *jwt_ECTokenService) Validate(token *jwt.Token) error {

	return ValidateToken(token)
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
