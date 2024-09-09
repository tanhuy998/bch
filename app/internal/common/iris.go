package common

import (
	"fmt"

	"github.com/kataras/iris/v12"
)

func SendDefaulJsonBodyAndEndRequest(ctx iris.Context, statusCode int, message string) {

	// body := auth_err_body_reponse{
	// 	Message: message,
	// }

	ctx.StatusCode(statusCode)
	ctx.JSON(fmt.Sprintf(`{message:%s}`, message))
	ctx.EndRequest()
}
