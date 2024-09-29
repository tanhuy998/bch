package port

import "github.com/kataras/iris/v12/mvc"

type (
	IActionResult interface {
		MarshallOutput(resultContent interface{}, response *mvc.Response) error
		NewActionResponse() *mvc.Response
	}
)
