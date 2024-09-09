package authServiceAdapter

type (
	AuthData struct {
	}

	IFetchAuthData interface {
		Serve(userUUID string) *AuthData
	}
)
