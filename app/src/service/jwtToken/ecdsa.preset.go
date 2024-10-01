package jwtTokenService

import (
	"crypto/ecdsa"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateECJWTToken(signingMethod *jwt.SigningMethodECDSA) *jwt.Token {

	//signingMethod := config.GetJWTEncryptionAlgo()

	token := jwt.New(signingMethod)

	return token
}

func SignECJWTToken(token *jwt.Token, privateKey *ecdsa.PrivateKey) (string, error) {

	return token.SignedString(privateKey)
}

func ValidateToken(token *jwt.Token) error {

	//claims := token.Claims

	return nil
}

func IsECSigningMethod(method interface{}) bool {

	_, ok := method.(*jwt.SigningMethodECDSA)

	return ok
}
