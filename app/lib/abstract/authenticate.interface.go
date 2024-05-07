package abstract

type IAuthenticate interface {
	VerifyPassword(string) error
	IsVerified() bool
}
