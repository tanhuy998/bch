package authService

import (
	"app/src/repository"
	"context"

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

// type AuthenticateService struct {
// 	vault *auth_vault
// }

type (
	// IIdentityManipulator interface {
	// 	CreateUser(username string, password string) error
	// 	AssignUserToCommandGroup(userUUID uuid.UUID, commandGroupUUID uuid.UUID) (*model.CommandGroupUser, error)
	// 	GrantCommandGroupUserRole(commandGroupUserUUID uuid.UUID, RoleUUID uuid.UUID) error
	// 	GetGroupMembers(GroupUUID uuid.UUID, pivot primitive.ObjectID, limit int, isPrev bool) (*repository.PaginationPack[authValueObject.CommandGroupUserEntity], error)
	// }

	IAuthorize interface {
		AuthorizeClaims(token string, field AuthorizationField, claims []AuthorizationGroup) (*context.Context, error)
		AuthorizeGroup(token string, field AuthorizationField, groups []AuthorizationGroup) (*context.Context, error)
	}

	IAuthenticate interface {
		ValidateCredential(uname string, pass string) (string, error)
		SignUp(uname string, pass string) error
		ChangePassword(uname string) error
	}

	IAuthService interface {
		IAuthenticate
		IAuthorize
	}

	AuthenticationService struct {
		UserRepo                 repository.IUser
		CommandGroupRepo         repository.ICommandGroup
		CommandGroupUserRepo     repository.ICommandGroupUser
		CommandGroupUserRoleRepo repository.ICommandGroupUserRole
		RoleRepo                 repository.IRole
	}
)

func (this *AuthenticationService) AuthorizeClaims(token string, field AuthorizationField, claims []AuthorizationGroup) (*context.Context, error) {

	return nil, nil
}

func (this *AuthenticationService) AuthorizeGroup(token string, field AuthorizationField, groups []AuthorizationGroup) (*context.Context, error) {

	return nil, nil
}

func (this *AuthenticationService) ValidateCredential(uname string, pass string) (string, error) {

	return "", nil
}

func (this *AuthenticationService) SignUp(uname string, pass string) error {

	return nil
}

func (this *AuthenticationService) ChangePassword(uname string) error {

	return nil
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
