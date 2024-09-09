package bootstrap

import (
	"crypto/ecdsa"
	"fmt"
	"os"
	"path"

	"github.com/golang-jwt/jwt/v5"
)

const (
	ENV_AUTH_JWT_PUBLIC_KEY_DIR  = "AUTH_JWT_PUBLIC_KEY_DIR"
	ENV_AUTH_JWT_PRIVATE_KEY_DIR = "AUTH_JWT_PRIVATE_KEY_DIR"
	//ENV_AUTH_JWT_ALG         = "AUTH_JWT_ALG"
)

var (
	jwt_private_key *ecdsa.PrivateKey
	jwt_public_key  *ecdsa.PublicKey
	jwt_algo        *jwt.SigningMethodECDSA = jwt.SigningMethodES256
)

func InitializeAuthEncryptionData() {

	fmt.Println("Reading auth key pair")

	__dir, err := os.Getwd()

	if err != nil {

		panic("could not retrieve working directory for reading auth keys")
	}

	var buffer []byte

	buffer, err = os.ReadFile(path.Join(__dir, os.Getenv(ENV_AUTH_JWT_PRIVATE_KEY_DIR)))

	if err != nil {

		panic("reading private key error: " + err.Error())
	}

	jwt_private_key, err = jwt.ParseECPrivateKeyFromPEM(buffer)

	if err != nil {

		panic("invalid private key")
	}

	buffer, err = os.ReadFile(path.Join(__dir, os.Getenv(ENV_AUTH_JWT_PUBLIC_KEY_DIR)))

	if err != nil {

		panic("reading public key error: " + err.Error())
	}

	jwt_public_key, err = jwt.ParseECPublicKeyFromPEM(buffer)

	if err != nil {

		panic("invalid public key")
	}

	jwt_private_key.PublicKey = *jwt_public_key
}

func GetJWTEncryptionPrivateKey() *ecdsa.PrivateKey {

	return jwt_private_key
}

func GetJWTEncryptionPublicKey() *ecdsa.PublicKey {

	return jwt_public_key
}

func GetJWTEncryptionAlgo() *jwt.SigningMethodECDSA {

	return jwt_algo
}
