package middleware

import (
	requestPresenter "app/domain/presenter/request"
	responsePresenter "app/domain/presenter/response"
	libError "app/lib/error"
	"io"

	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/hero"
)

type (
	PresenterInitializer[RequestPresenter_T, ResponsePresenter_T any] func(req *RequestPresenter_T, res *ResponsePresenter_T)
)

func BindPresenters[RequestPresenter_T any, ResponsePresenter_T any](
	container *hero.Container,
	initializers ...PresenterInitializer[RequestPresenter_T, ResponsePresenter_T],
) iris.Handler {

	if container == nil {

		panic("BindPresenter middleware need container to function")
	}

	return container.Handler(func(ctx iris.Context, validator context.Validator) {

		if validator == nil {

			ctx.StopWithJSON(500, &responsePresenter.ErrorResponse{
				Message: "no validator",
			})
			return
		}

		var (
			request  *RequestPresenter_T  = new(RequestPresenter_T)
			response *ResponsePresenter_T = new(ResponsePresenter_T)
			err      error
		)

		runInitializers(request, response, initializers)

		if p, ok := any(request).(requestPresenter.IRequestBinder); ok {

			err = p.Bind(ctx)

		} else {

			err = bindDefault(request, ctx)
		}

		if !libError.IsAcceptable(err, io.EOF) {
			/*
				io.EOF returned when request body is empty
			*/
			ctx.StopWithJSON(400, &responsePresenter.ErrorResponse{
				Message: err.Error(),
			})
			return
		}

		err = validator.Struct(request)

		if err != nil {

			ctx.StopWithJSON(400, &responsePresenter.ErrorResponse{
				Message: err.Error(),
			})
			return
		}

		ctx.RegisterDependency(request)
		ctx.RegisterDependency(response)
		ctx.Next()
	})
}

func runInitializers[RequestPresenter_T, ResponsePresenter_T any](
	req *RequestPresenter_T,
	res *ResponsePresenter_T,
	initializers []PresenterInitializer[RequestPresenter_T, ResponsePresenter_T],
) {

	for _, f := range initializers {

		f(req, res)
	}
}

func bindDefault[RequestPresenter_T any](presenter *RequestPresenter_T, ctx iris.Context) error {

	ctx.ReadURL(presenter)
	ctx.ReadJSON(presenter)

	return nil
}

func isValidationError(err error) bool {

	if _, ok := err.(*validator.InvalidValidationError); ok {

		return true
	}

	if _, ok := err.(validator.ValidationErrors); ok {

		return true
	}

	return false
}
