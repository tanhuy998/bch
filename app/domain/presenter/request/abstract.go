package requestPresenter

import "github.com/kataras/iris/v12"

type IRequestBinder interface {
	Bind(ctx iris.Context)
}
