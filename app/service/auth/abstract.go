package authService

type IAuthorize interface {
	AuthorizeClaims(token string, field string, claims []string) error
	AuthorizeGroup(token string, field string, groups []string) error
}

type IAuthenticate interface {
	ValidateCredential(uname string, pass string) (string, error)
}

type IAuthCore interface {
	IAuthenticate
	IAuthorize
}
