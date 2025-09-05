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
	PresenterBindingHook[RequestPresenter_T, ResponsePresenter_T any] func(container *hero.Container, ctx iris.Context, req *RequestPresenter_T, res *ResponsePresenter_T) error
	RequestPresenterInitializer[RequestPresenter_T any]               func(req *RequestPresenter_T) error
)

func UseAuthority[Req_T requestInput.IAuthorityBringAlong, Res_T any](
	container *hero.Container, ctx iris.Context, req Req_T, res Res_T,
) error {

	// accessToken := libIris.GetAccessToken(ctx)

	// if accessToken == nil {

	// 	container.Handler(readAccessToken)(ctx)

	// 	accessToken = libIris.GetAccessToken(ctx)
	// }

	// if accessToken == nil {

	// 	return nil
	// }

	// req.SetAuthority(accessToken.GetAuthData())
	// return nil

	return binding.UseAuthority(container, ctx, req, res)
}

func UseTenantMapping[Req_T requestInput.ITenantMappingInput, Res_T any](
	container *hero.Container, ctx iris.Context, req Req_T, res Res_T,
) error {

	// accessToken := libIris.GetAccessToken(ctx)

	// if accessToken == nil {

	// 	container.Handler(readAccessToken)(ctx)

	// 	accessToken = libIris.GetAccessToken(ctx)
	// }

	// if accessToken == nil {

	// 	return nil
	// }

	// req.SetTenantUUID(accessToken.GetTenantUUID())
	// return nil

	return binding.UseTenantMapping(container, ctx, req, res)
}

// func readAccessToken(
// 	ctx iris.Context, accessTokenClient accessTokenClientPort.IAccessTokenClient,
// ) error {

// 	at, err := accessTokenClient.Read(ctx)

// 	if err != nil {

// 		return err
// 	}

// 	libIris.SetAccessToken(ctx, at)

// 	return nil
// }
