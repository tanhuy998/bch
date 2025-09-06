package jwt

import (
	"app/internal/bootstrap"
	jwtTokenService "app/service/jwtToken"

	"github.com/golang-jwt/jwt/v5"
)

// var (
// 	AsymmetricJWTService = jwtTokenService.NewECDSAService(
// 		//jwt.SigningMethodES256, *GetJWTAsymmetricEncryptionPrivateKey(), *GetJWTAsymmetricEncryptionPublicKey(),
// 		jwt.SigningMethodES256, *jwt_asym_private_key, *jwt_asym_public_key,
// 	)

// 	SymmetricJWTService = jwtTokenService.NewHMACService(
// 		// jwt.SigningMethodHS256, bootstrap.GetJWTSymmetricEncryptionSecret(),
// 		jwt.SigningMethodHS256, jwt_symmetric_secret,
// 	)
// )

func NewAsymetricJWTService() *jwtTokenService.JWTECTokenService {

	return jwtTokenService.NewECDSAService(
		jwt.SigningMethodES256, *bootstrap.GetJWTAsymmetricEncryptionPrivateKey(), *bootstrap.GetJWTAsymmetricEncryptionPublicKey(),
		// jwt.SigningMethodES256, *jwt_asym_private_key, *jwt_asym_public_key,
	)
}

func NewSymetricJWTService() *jwtTokenService.JWTHMACTokenService {

	return jwtTokenService.NewHMACService(
		jwt.SigningMethodHS256, bootstrap.GetJWTSymmetricEncryptionSecret(),
		// jwt.SigningMethodHS256, jwt_symmetric_secret,
	)
}
