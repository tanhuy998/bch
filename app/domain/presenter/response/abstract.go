package responsePresenter

import (
	"github.com/kataras/iris/v12"
)

type IResponsePresenter interface {
	Bind(ctx iris.Context) error
}

type PaginationNavigation struct {
	Previous string `json:"previous"`
	Next     string `json:"next"`
}
