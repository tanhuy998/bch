package authService

import "context"

type IAuthorize interface {
	AuthorizeClaims(token string, field AuthorizationField, claims []AuthorizationGroup) (*context.Context, error)
	AuthorizeGroup(token string, field AuthorizationField, groups []AuthorizationGroup) (*context.Context, error)
}

type IAuthenticate interface {
	ValidateCredential(uname string, pass string) (string, error)
	SignUp(uname string, pass string) error
	ChangePassword(uname string) error
}

type IAuthService interface {
	IAuthenticate
	IAuthorize
}
