package abstract

type Role string

type IAuthorize interface {
	GetRoles()
	Authorize([]Role) error
}
