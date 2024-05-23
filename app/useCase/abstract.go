package usecase

import (
	"encoding/json"

	"github.com/kataras/iris/v12/mvc"
)

type (
	IResponseWrapper interface {
	}

	ResponseWrapper[T any] struct {
		Target T
	}

	ActionResultUseCase struct {
	}
)

func (this *ActionResultUseCase) MarshallOutput(resultContent interface{}, response *mvc.Response) error {

	res := NewResponse()

	return MarshalResponseContent(resultContent, res)
}

func (this *ActionResultUseCase) NewActionResponse() *mvc.Response {

	return NewResponse()
}

func NewResponse() *mvc.Response {

	return &mvc.Response{
		ContentType: "application/json",
	}
}

func MarshalResponseContent(context interface{}, res *mvc.Response) error {

	data, err := json.Marshal(context)

	if err != nil {

		return err
	}

	res.Content = data

	return nil
}
