package jwtTokenService

import (
	"crypto/ecdsa"

	"github.com/golang-jwt/jwt/v5"
)

var (
	signing_method = jwt.SigningMethodES256
)

type (
	jwt_ECTokenService struct {
		private_key *ecdsa.PrivateKey
		public_key  *ecdsa.PublicKey
		//signing_method *jwt.SigningMethodECDSA
	}
)

func NewECDSAService(privateKey ecdsa.PrivateKey, publicKey ecdsa.PublicKey) *jwt_ECTokenService {

	ret := &jwt_ECTokenService{
		private_key: &privateKey,
		public_key:  &publicKey,
	}

	return ret
}

func (this jwt_ECTokenService) GenerateToken() *jwt.Token {

	return GenerateECJWTToken(signing_method)
}

func (this jwt_ECTokenService) SignedString(token *jwt.Token) (string, error) {

	return token.SignedString(this.private_key)
}

func (this jwt_ECTokenService) VerifyTokenString(token_str string) (*jwt.Token, error) {

	return jwt.Parse(token_str, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {

			return nil, ERR_SIGNING_METHOD_MISMATCH
		}

		return this.public_key, nil
	})
}

func (this jwt_ECTokenService) Validate(token *jwt.Token) error {

	return ValidateToken(token)
}
