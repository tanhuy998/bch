package actionResultService

import (
	actionResultServicePort "app/port/actionResult"
	"encoding/json"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type (
	// IActionResult interface {
	// 	Prepare() IResponse
	// 	ServeResponse(content interface{}) (IResponse, error)
	// 	ServeErrorResponse(error) (IResponse, error)
	// }

	// IResponse interface {
	// 	mvc.Result
	// 	SetCode(int) IResponse
	// 	Redirect(link string) IResponse
	// 	ContentType(string) IResponse
	// 	SetContent([]byte) IResponse
	// 	Done() (IResponse, error)
	// 	ServeResponse(content interface{}) (IResponse, error)
	// 	ServeErrorResponse(error) (IResponse, error)
	// }

	/*
		ResponseResultService defines preset for application/json content type
		for Iris result.
	*/
	ResponseResultService struct{}

	Iris_Response = mvc.Response

	ResponseEntity struct {
		*Iris_Response
	}

	ErrorResponseFormat struct {
		Message string `json:"message"`
	}
)

func _ServeResponse(res actionResultServicePort.IResponse, content interface{}) (actionResultServicePort.IResponse, error) {

	bytes, err := json.Marshal(content)

	if err != nil {

		return nil, err
	}

	res.SetContent(bytes)
	res.SetCode(200)

	return res, nil
}

func _ServeErrorResponse(res actionResultServicePort.IResponse, err error) (actionResultServicePort.IResponse, error) {

	content := &ErrorResponseFormat{
		Message: err.Error(),
	}

	bytes, err := json.Marshal(content)

	if err != nil {

		return nil, err
	}

	res.SetCode(400)
	res.SetContent(bytes)

	return res, nil
}

func (this *ResponseResultService) Prepare() actionResultServicePort.IResponse {

	return NewResponseEntity()
}

func (this *ResponseResultService) ServeResponse(content interface{}) (actionResultServicePort.IResponse, error) {

	return _ServeResponse(NewResponseEntity(), content)
}

func (this *ResponseResultService) ServeErrorResponse(err error) (actionResultServicePort.IResponse, error) {

	return _ServeErrorResponse(NewResponseEntity(), err)
}

func NewResponseEntity() actionResultServicePort.IResponse {

	return &ResponseEntity{
		Iris_Response: &mvc.Response{
			ContentType: "application/json",
		},
	}
}

func (this *ResponseEntity) Dispatch(ctx iris.Context) {

	this.Iris_Response.Dispatch(ctx)
}

func (this *ResponseEntity) SetCode(code int) actionResultServicePort.IResponse {

	this.Iris_Response.Code = code

	return this
}

func (this *ResponseEntity) Redirect(link string) actionResultServicePort.IResponse {

	this.Iris_Response.Path = link

	return this
}

func (this *ResponseEntity) ContentType(t string) actionResultServicePort.IResponse {

	this.Iris_Response.ContentType = t

	return this
}

func (this *ResponseEntity) SetContent(content []byte) actionResultServicePort.IResponse {

	this.Iris_Response.Content = content

	return this
}

func (this *ResponseEntity) ServeResponse(content interface{}) (actionResultServicePort.IResponse, error) {

	return _ServeResponse(this, content)
}

func (this *ResponseEntity) ServeErrorResponse(err error) (actionResultServicePort.IResponse, error) {

	return _ServeErrorResponse(this, err)
}

func (this *ResponseEntity) Done() (actionResultServicePort.IResponse, error) {

	return this, nil
}
