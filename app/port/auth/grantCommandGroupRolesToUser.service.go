package authServicePort

type (
	IGrantCommandGroupRolesToUser interface {
		Serve(groupUUID string, userUUID string, roles []string) error
	}
)
