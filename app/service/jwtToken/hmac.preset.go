package jwtTokenService

import (
	"github.com/golang-jwt/jwt/v5"
)

func GenerateHMACToken(signingMethod *jwt.SigningMethodHMAC) *jwt.Token {

	//signingMethod := config.GetJWTEncryptionAlgo()

	token := jwt.New(signingMethod)

	return token
}

func SignHMACToken(token *jwt.Token, secret []byte) (string, error) {

	return token.SignedString(secret)
}

func ValidateHMACToken(token *jwt.Token) error {

	//claims := token.Claims

	return nil
}

func IsHMACSigningMethod(method interface{}) bool {

	_, ok := method.(*jwt.SigningMethodHMAC)

	return ok
}
