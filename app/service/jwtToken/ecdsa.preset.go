package jwtTokenService

import (
	"crypto/ecdsa"
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ERR_SIGNING_METHOD_MISMATCH = errors.New("jwt singing method mismatch")
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
