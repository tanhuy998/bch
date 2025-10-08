package hook

import (
	"app/infrastructure/restfull/common/middleware/hook/binding"
	"app/valueObject/requestInput"

	"errors"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
)

var (
	ERR_NO_CONTEXT = errors.New("bindPresenters helper: no context")
)

type (
	PresenterBindingHook[RequestPresenter_T, ResponsePresenter_T any] binding.Hook
	RequestPresenterInitializer[RequestPresenter_T any]               func(req *RequestPresenter_T) error
)

func UseAuthority[Req_T requestInput.IAuthorityBringAlong, Res_T any](
	container *hero.Container, ctx iris.Context, req Req_T, res Res_T,
) error {

	return binding.UseAuthority(container, ctx, req, res)
}

func UseTenantMapping[Req_T requestInput.ITenantMappingInput, Res_T any](
	container *hero.Container, ctx iris.Context, req Req_T, res Res_T,
) error {

	return binding.UseTenantMapping(container, ctx, req, res)
}
