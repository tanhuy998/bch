package authServicePort

type (
	IAddUserToCommandGroup interface {
		Serve(groupUUID string, userUUID string) error
		Get() IGetSingleCommandGroup
	}
)
