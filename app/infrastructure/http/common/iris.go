package common

import (
	accessTokenServicePort "app/port/accessToken"

	"github.com/kataras/iris/v12"
)

type (
	reponse_body struct {
		Message string `json:"message"`
	}
)

func SendDefaulJsonBodyAndEndRequest(ctx iris.Context, statusCode int, message string) {

	ctx.StatusCode(statusCode)
	ctx.JSON(reponse_body{
		Message: message,
	})
	ctx.EndRequest()
}

func GetAccessToken(ctx iris.Context) accessTokenServicePort.IAccessToken {

	unknown := ctx.Values().Get(CTX_ACCESS_TOKEN_KEY)

	if accessToken, ok := unknown.(accessTokenServicePort.IAccessToken); ok {

		return accessToken
	}

	return nil
}
