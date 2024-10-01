package actionResultServicePort

import "github.com/kataras/iris/v12/mvc"

type (
	// IActionResult interface {
	// 	MarshallOutput(resultContent interface{}, response *mvc.Response) error
	// 	NewActionResponse() *mvc.Response
	// }

	IActionResult interface {
		Prepare() IResponse
		ServeResponse(content interface{}) (IResponse, error)
		ServeErrorResponse(error) (IResponse, error)
	}

	IResponse interface {
		mvc.Result
		SetCode(int) IResponse
		Redirect(link string) IResponse
		ContentType(string) IResponse
		SetContent([]byte) IResponse
		Done() (IResponse, error)
		ServeResponse(content interface{}) (IResponse, error)
		ServeErrorResponse(error) (IResponse, error)
	}
)
