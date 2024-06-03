package actionResultService

import (
	"encoding/json"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type (
	IActionResult interface {
		ServeResponse(content interface{}) (IResponseBuilder, error)
		ServeErrorResponse(error) (IResponseBuilder, error)
	}

	IResponseBuilder interface {
		mvc.Result
		SetCode(int) IResponseBuilder
		Redirect(link string) IResponseBuilder
		ContentType(string) IResponseBuilder
		SetContent([]byte) IResponseBuilder
	}

	/*
		ResponseResultService defines preset for application/json content type
		for Iris result.
	*/
	ResponseResultService struct {
	}

	Iris_Response = mvc.Response

	ResponseEntity struct {
		*Iris_Response
	}

	ErrorResponseFormat struct {
		Message string `json:"message"`
	}
)

func (this *ResponseResultService) ServeResponse(content interface{}) (IResponseBuilder, error) {

	res := NewResponseEntity()

	bytes, err := json.Marshal(content)

	if err != nil {

		return nil, err
	}

	res.SetContent(bytes)
	res.SetCode(200)

	return res, nil
}

func (this *ResponseResultService) ServeErrorResponse(err error) (IResponseBuilder, error) {

	res := NewResponseEntity()

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

func NewResponseEntity() IResponseBuilder {

	return &ResponseEntity{
		Iris_Response: &mvc.Response{
			ContentType: "application/json",
		},
	}
}

func (this *ResponseEntity) Dispatch(ctx iris.Context) {

	this.Iris_Response.Dispatch(ctx)
}

func (this *ResponseEntity) SetCode(code int) IResponseBuilder {

	this.Iris_Response.Code = code

	return this
}

func (this *ResponseEntity) Redirect(link string) IResponseBuilder {

	this.Iris_Response.Path = link

	return this
}

func (this *ResponseEntity) ContentType(t string) IResponseBuilder {

	this.Iris_Response.ContentType = t

	return this
}

func (this *ResponseEntity) SetContent(content []byte) IResponseBuilder {

	this.Iris_Response.Content = content

	return this
}
