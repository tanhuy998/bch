package accessTokenServicePort

type (
	IAccessTokenBringAlong interface {
		//IContextBringAlong
		ReceiveAccessToken(at IAccessToken)
		GetAccessToken() IAccessToken
	}
)
