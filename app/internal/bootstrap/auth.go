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
	ENV_APP_SECRET               = "APP_SECRET"
	//ENV_AUTH_JWT_ALG         = "AUTH_JWT_ALG"
)

var (
	jwt_asym_private_key *ecdsa.PrivateKey
	jwt_asym_public_key  *ecdsa.PublicKey

	jwt_symmetric_secret []byte
	jwt_algo             *jwt.SigningMethodECDSA = jwt.SigningMethodES256
)

func initializeAuthEncryptionData() {

	fmt.Println("Reading auth key pair")

	__dir, err := os.Getwd()

	if err != nil {

		panic("could not retrieve working directory for reading auth keys")
	}

	var buffer []byte

	buffer, err = os.ReadFile(path.Join(__dir, os.Getenv(ENV_AUTH_JWT_PRIVATE_KEY_DIR)))

	if err != nil {

		panic(fmt.Sprintf("reading private key error: %s", err.Error()))
	}

	jwt_asym_private_key, err = jwt.ParseECPrivateKeyFromPEM(buffer)

	if err != nil {

		panic(fmt.Sprintf("private key error: %s", err.Error()))
	}

	buffer, err = os.ReadFile(path.Join(__dir, os.Getenv(ENV_AUTH_JWT_PUBLIC_KEY_DIR)))

	if err != nil {

		panic(fmt.Sprintf("reading public key error: %s", err.Error()))
	}

	jwt_asym_public_key, err = jwt.ParseECPublicKeyFromPEM(buffer)

	if err != nil {

		panic(fmt.Sprintf("public key error: %s", err.Error()))
	}

	jwt_asym_private_key.PublicKey = *jwt_asym_public_key

	sym_secret_str := os.Getenv(ENV_APP_SECRET)
	jwt_symmetric_secret = []byte(sym_secret_str)
}

func GetJWTSymmetricEncryptionSecret() []byte {

	ret := make([]byte, len(jwt_symmetric_secret))
	copy(ret, jwt_symmetric_secret)

	return ret
}

func GetJWTAsymmetricEncryptionPrivateKey() *ecdsa.PrivateKey {

	return jwt_asym_private_key
}

func GetJWTAsymmetricEncryptionPublicKey() *ecdsa.PublicKey {

	return jwt_asym_public_key
}

func GetJWTEncryptionAlgo() *jwt.SigningMethodECDSA {

	return jwt_algo
}
