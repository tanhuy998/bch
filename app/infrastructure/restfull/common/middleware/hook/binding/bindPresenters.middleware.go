package binding

import (
	libCommon "app/internal/lib/common"
	libIris "app/internal/lib/iris"
	"app/shared/common/object"
	"app/valueObject/requestInput"

	"io"
	"net/http"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/hero"
)

type IRequestBinder interface {
	/*
		Structs that implement IRequestBinder define its own
		request context to add extra business logic after data
		transfered from request context to the request presenter object.
	*/
	Bind(ctx iris.Context) error
}

type ICustomQueryParamsBinder interface {
	BindCustomQueryParams(queryParams map[string]string) error
}

type (
	IQueryParamsBinder interface {
		GetBinderReference() interface{}
	}
)

type (
	ContainerDependentMiddleware func(container *hero.Container) context.Handler
)

type EmptyPresenter struct {
}

type (
	BindingHandler[RequestPresenter_T any, ResponsePresenter_T any] struct {
		Validator    context.Validator
		requestPool  object.ObjectRecycler[RequestPresenter_T]
		responsePool object.ObjectRecycler[ResponsePresenter_T]
	}
)

func resolveoObject[T any](pool *object.ObjectRecycler[T]) *T {

	if libCommon.IsInterface[T]() {

		panic("presenter must be type of concrete, not abstract")
	}

	if isEmptyPresenter[T]() {

		return nil
	}

	switch obj := pool.Load(); {
	case obj == nil:
		obj = new(T)
		return obj
	default:
		return obj
	}
}

func (this *BindingHandler[RequestPresenter_T, ResponsePresenter_T]) collectPresenters(req *RequestPresenter_T, res *ResponsePresenter_T) {

	if req != nil {

		this.requestPool.Collect(req)
	}

	if res != nil {

		this.responsePool.Collect(res)
	}
}

func (this *BindingHandler[RequestPresenter_T, ResponsePresenter_T]) Handle(container *hero.Container, ctx iris.Context, hooks []Hook) {

	if this.Validator == nil {

		libIris.SendDefaulJsonBodyAndEndRequest(ctx, http.StatusInternalServerError, "no validator")
		return
	}

	var (
		request  *RequestPresenter_T  = resolveoObject(&this.requestPool)
		response *ResponsePresenter_T = resolveoObject(&this.responsePool)
		err      error
	)

	if val, ok := any(request).(requestInput.IContextBringAlong); ok {

		val.ReceiveContext(ctx)
	}

	err = runInitializers(container, ctx, request, response, hooks)

	if err != nil {

		libIris.SendDefaulJsonBodyAndEndRequest(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if p, ok := any(request).(IRequestBinder); ok {

		err = p.Bind(ctx)

	} else {

		err = bindRequestDefault(request, ctx)
	}

	switch err {
	case nil:
	case io.EOF:
	default:
		libIris.SendDefaulJsonBodyAndEndRequest(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := this.Validator.Struct(request); err != nil {

		libIris.SendDefaulJsonBodyAndEndRequest(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if request != nil {

		ctx.RegisterDependency(request)
	}

	if response != nil {

		ctx.RegisterDependency(response)
	}

	ctx.Next()
	this.collectPresenters(request, response)
}

func BindPresenters[RequestPresenter_T any, ResponsePresenter_T any](
	hooks ...Hook,
) ContainerDependentMiddleware {

	return func(container *hero.Container) context.Handler {

		if container == nil {

			panic("BindPresenter middleware need container to function")
		}

		ensureProperPresenteTypes[RequestPresenter_T, ResponsePresenter_T]()

		// handlerObject := new(BindingHandler[RequestPresenter_T, ResponsePresenter_T])

		container.EnableStructDependents = true

		container.Register(
			new(BindingHandler[RequestPresenter_T, ResponsePresenter_T]),
		).Explicitly().EnableStructDependents()

		// return func(ctx iris.Context) {

		// 	handlerObject.Handle(container, ctx, hooks)
		// }

		return container.Handler(
			func(ctx iris.Context, handlerObj *BindingHandler[RequestPresenter_T, ResponsePresenter_T]) {

				handlerObj.Handle(container, ctx, hooks)
			},
		)
	}
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
	case *EmptyPresenter:
		return true
	default:
		return false
	}
}

func runInitializers[RequestPresenter_T, ResponsePresenter_T any](
	container *hero.Container,
	ctx iris.Context,
	req *RequestPresenter_T,
	res *ResponsePresenter_T,
	initializers []Hook,
) error {

	for _, f := range initializers {

		err := f(container, ctx, req, res)

		if err != nil {

			return err
		}
	}

	return nil
}

func bindRequestDefault[RequestPresenter_T any](presenter *RequestPresenter_T, ctx iris.Context) error {

	ctx.ReadURL(presenter)
	ctx.ReadJSON(presenter)

	switch v, ok := any(presenter).(IQueryParamsBinder); {
	case ok:

		binderReference := v.GetBinderReference()

		// requestQuery := ctx.Request().URL.Query()

		// bytes, _ := json.Marshal(requestQuery)

		// if len(bytes) == 0 {

		// 	return nil
		// }

		// return json.Unmarshal(bytes, binderReference)

		return ctx.ReadQuery(binderReference)
	}

	return nil
}
