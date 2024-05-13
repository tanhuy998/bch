package authService

import (
	"fmt"

	"github.com/gofor-little/env"
	"github.com/golang-jwt/jwt/v5"
)

const ENV_JWT_PRIVATE_KEY = "JWT_PRIVATE_KEY"

var (
	private_key    []byte
	signing_method *jwt.SigningMethodECDSA = jwt.SigningMethodES256
)

/*
Authentication service is design in microservice manner in order to isolate the authentication
and authorization logics with the app's bussiness logic
in inplementing Policy Decision Point pattern.

Authorization implement the following principal
claim-based authorization
group-based authorization

This service is determined as a centralized authentication/authorization like an
authentication server for internal authentication/authorization between services in a distributed
system. It manipulates it's own authentication database for authentication
and authorization. Applications that use this service will be provided an AppID in order to
indentify its right to use and interact with the service. Besides, applications also been
provided a public key to distributedly verify user's access tokens.

There are 4 concept
+ Field is the target object that groups aim to be granted to
+ Group is a set of users that will is organized to be gramted a specific permission
+ License is the permission on a field that a group could manifest
+ Claim is the info about the license, sometimes it also been detemined as the operation
that users of a granted goup could take action.
*/

type AuthorizationClaim string
type AuthorizationGroup string
type AuthorizationField string

type AuthorizationLicense struct {
	Fields AuthorizationField
	Groups []AuthorizationGroup
	Claims []AuthorizationClaim
}

type AuthenticateService struct {
	vault *auth_vault
}

func New(connString string) *AuthenticateService {

	ret := new(AuthenticateService)
	ret.vault = &auth_vault{}

	ret.SetConnString(connString)
	privateKey, err := retrievePrivateKey()

	if err != nil {

		ret.vault.pending_error = err
	}

	ret.vault.private_key = privateKey

	return ret
}

func (this *AuthenticateService) SetConnString(connString string) {

	this.vault.conn_string = connString
}

// func (this *AuthenticateService) ValidateCredential(uname string, pass string) (string, error) {

// 	//return this.core.ValidateCredential(uname, pass)
// }

// func (this *AuthenticateService) AuthorizeClaims(token string, field string, claims []string) error {

// 	//return this.core.AuthorizeClaims(token, field, claims)
// }

// func (this *AuthenticateService) AuthorizeGroup(token string, field string, groups []string) error {

// 	//return this.core.AuthorizeGroup(token, field, groups)
// }

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
