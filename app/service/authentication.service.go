package service

import (
	"app/app/model"
	"fmt"

	"github.com/gofor-little/env"
	"go.mongodb.org/mongo-driver/mongo"
)

const ENV_JWT_PRIVATE_KEY = "JWT_PRIVATE_KEY"

var private_key []byte

/*
Authentication service is design in microservice manner in order to isolate the authentication
and authorization logics with the app's bussiness logic
in inplementing Policy Decision Point pattern.

Authorization implement the following principal
claim-based authorization
group-based authorization
document-based authorization
*/

type AuthenticateService struct {
	auth_db *mongo.Database
}

func (this *AuthenticateService) SetDB(db *mongo.Database) {

	this.auth_db = db
}

func (this *AuthenticateService) ValidateCredential(inputUser model.User) (string, error) {

	return "", nil
}

func (this *AuthenticateService) Authorize(token string)

func generateToken(user model.User) string {

	return ""
}

func retrievePrivateKey() ([]byte, error) {

	if len(private_key) > 0 {

		return private_key, nil
	}

	env := env.Get(ENV_JWT_PRIVATE_KEY, "")

	if env == "" {

		return nil, fmt.Errorf("")
	}

	private_key = []byte(env)

	return private_key, nil
}
