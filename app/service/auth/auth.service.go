package authService

import (
	"fmt"

	"github.com/gofor-little/env"
	"github.com/golang-jwt/jwt/v5"
)

const ENV_JWT_PRIVATE_KEY = "JWT_PRIVATE_KEY"

var private_key []byte
var signing_method *jwt.SigningMethodECDSA = jwt.SigningMethodES256

/*
Authentication service is design in microservice manner in order to isolate the authentication
and authorization logics with the app's bussiness logic
in inplementing Policy Decision Point pattern.

Authorization implement the following principal
claim-based authorization
group-based authorization
document-based authorization
*/

type AuthorizationClaim string
type AuthorizationGroup string
type AuthorizationField string

type AuthenticateService struct {
	IAuthCore
	core *auth_core
}

func New(connString string) *AuthenticateService {

	ret := new(AuthenticateService)
	ret.core = &auth_core{}

	ret.SetConnString(connString)
	privateKey, err := retrievePrivateKey()

	if err != nil {

		ret.core.pending_error = err
	}

	ret.core.private_key = privateKey

	return ret
}

func (this *AuthenticateService) SetConnString(connString string) {

	this.core.conn_string = connString
}

func (this *AuthenticateService) ValidateCredential(uname string, pass string) (string, error) {

	return this.core.ValidateCredential(uname, pass)
}

func (this *AuthenticateService) AuthorizeClaims(token string, field string, claims []string) error {

	return this.core.AuthorizeClaims(token, field, claims)
}

func (this *AuthenticateService) AuthorizeGroup(token string, field string, groups []string) error {

	return this.core.AuthorizeGroup(token, field, groups)
}

func retrievePrivateKey() ([]byte, error) {

	if len(private_key) > 0 {

		return private_key, nil
	}

	env := env.Get(ENV_JWT_PRIVATE_KEY, "")

	if env == "" {

		return nil, fmt.Errorf("No Private key for authentication")
	}

	private_key = []byte(env)

	return private_key, nil
}
