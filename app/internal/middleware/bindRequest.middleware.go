package middleware

import (
	requestPresenter "app/domain/presenter/request"
	libError "app/lib/error"
	"fmt"
	"io"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

func BindRequest[RequestPresenter_T any]() iris.Handler {

	return func(ctx iris.Context) {

		var (
			presenter *RequestPresenter_T = new(RequestPresenter_T)
			err       error
		)

		if p, ok := any(presenter).(requestPresenter.IRequestBinder); ok {

			err = p.Bind(ctx)

		} else {

			err = bindDefault(presenter, ctx)
		}

		if libError.IsAcceptable(err, io.EOF) {
			/*
				io.EOF is just signal that indicates the read operation reaches eof,
				not an impact error.
			*/
			ctx.RegisterDependency(presenter)
		}

		ctx.Next()
	}
}

func bindDefault[RequestPresenter_T any](presenter *RequestPresenter_T, ctx iris.Context) error {
	fmt.Println((ctx.Params().Len()), ctx.Params())
	err := ctx.ReadParams(presenter)
	fmt.Println("a")
	if err != nil {

		return err
	}

	err = ctx.ReadQuery(presenter)
	fmt.Println("b")
	if err != nil {

		return err
	}

	err = ctx.ReadJSON(presenter, context.JSONReader{
		DisallowUnknownFields: true,
	})
	fmt.Println("c", err)
	return err
}
