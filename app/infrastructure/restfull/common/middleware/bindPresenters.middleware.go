package middleware

import (
	"app/infrastructure/restfull/common/middleware/hook/binding"

	"github.com/kataras/iris/v12"
)

func init() {

	// schema.Query.ZeroEmpty(true)
}

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

type EmptyPresenter struct {
}

func BindPresenters[RequestPresenter_T any, ResponsePresenter_T any](
	hooks ...binding.Hook,
) ContainerDependentMiddleware {

	return binding.BindPresenters[RequestPresenter_T, ResponsePresenter_T](hooks...)
}
