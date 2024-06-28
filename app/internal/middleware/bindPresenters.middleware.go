package middleware

import (
	"app/domain/presenter"
	requestPresenter "app/domain/presenter/request"
	responsePresenter "app/domain/presenter/response"
	libCommon "app/lib/common"
	libError "app/lib/error"
	"io"

	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/hero"
)

type (
	PresenterInitializer[RequestPresenter_T, ResponsePresenter_T any] func(req *RequestPresenter_T, res *ResponsePresenter_T)
	RequestPresenterInitializer[RequestPresenter_T any]               func(req *RequestPresenter_T)
)

func BindPresenters[RequestPresenter_T any, ResponsePresenter_T any](
	container *hero.Container,
	initializers ...PresenterInitializer[RequestPresenter_T, ResponsePresenter_T],
) iris.Handler {

	if container == nil {

		panic("BindPresenter middleware need container to function")
	}

	ensureProperPresenteTypes[RequestPresenter_T, ResponsePresenter_T]()

	return container.Handler(func(ctx iris.Context, validator context.Validator) {

		if validator == nil {

			ctx.StopWithJSON(500, &responsePresenter.ErrorResponse{
				Message: "no validator",
			})
			return
		}

		var (
			request  *RequestPresenter_T  = instantiatePresenter[RequestPresenter_T]()  //new(RequestPresenter_T)
			response *ResponsePresenter_T = instantiatePresenter[ResponsePresenter_T]() //new(ResponsePresenter_T)
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

		if request != nil {

			ctx.RegisterDependency(request)
		}

		if response != nil {

			ctx.RegisterDependency(response)
		}

		ctx.Next()
	})
}

func ensureProperPresenteTypes[RequestPresenter_T, ResponsePresenter_T any]() {

	if !libCommon.IsInterface[RequestPresenter_T]() ||
		isEmptyPresenter[RequestPresenter_T]() &&
			!libCommon.IsInterface[ResponsePresenter_T]() ||
		isEmptyPresenter[ResponsePresenter_T]() {

		return
	}

	panic("could not bind types that is not struct type or IEmptyPresenter")
}

func instantiatePresenter[Presenter_T any]() *Presenter_T {

	if !libCommon.IsInterface[Presenter_T]() {

		return new(Presenter_T)
	}

	if isEmptyPresenter[Presenter_T]() {

		return nil
	}

	panic("Could not bind interface type as concrete presenter.")
}

func isEmptyPresenter[T any]() bool {

	var r any = (*T)(nil)

	switch r.(type) {
	case *presenter.IEmptyPresenter:
		return true
	default:
		return false
	}
}

// func BindRequestPresenter[RequestPresenter_T any](
// 	container *hero.Container,
// 	initilaizers ...RequestPresenterInitializer[RequestPresenter_T],
// ) {

// 	mustHaveContainer(container)

// 	return container.Handler(func(ctx iris.Context, validator context.Validator) {

// 	})
// }

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

func mustHaveContainer(container *hero.Container) {

	if container == nil {

		panic("BindPresenter middleware need container to function")
	}
}
