package authService

import (
	"fmt"
	"os"
)

func retrievePrivateKey() ([]byte, error) {

	if len(private_key) > 0 {

		return private_key, nil
	}

	env := os.Getenv(ENV_JWT_PRIVATE_KEY)

	if env == "" {

		return nil, fmt.Errorf("No Private key for authentication")
	}

	private_key = []byte(env)

	return private_key, nil
}
