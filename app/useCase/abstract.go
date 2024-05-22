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
)

func newResponse() *mvc.Response {

	return &mvc.Response{
		ContentType: "application/json",
	}
}

func marshalResponseContent(context interface{}, res *mvc.Response) error {

	data, err := json.Marshal(context)

	if err != nil {

		return err
	}

	res.Content = data

	return nil
}
