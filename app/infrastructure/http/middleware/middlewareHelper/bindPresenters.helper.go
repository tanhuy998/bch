package middlewareHelper

import (
	"app/infrastructure/http/common"
	accessTokenServicePort "app/port/accessToken"
	"app/valueObject/requestInput"

	"errors"

	"github.com/kataras/iris/v12"
)

var (
	ERR_NO_CONTEXT = errors.New("bindPresenters helper: no context")
)

type (
	PresenterInitializer[RequestPresenter_T, ResponsePresenter_T any] func(ctx iris.Context, req *RequestPresenter_T, res *ResponsePresenter_T) error
	RequestPresenterInitializer[RequestPresenter_T any]               func(req *RequestPresenter_T) error
)

func UseAccessToken[Req_T accessTokenServicePort.IAccessTokenBringAlong, Res_T any](ctx iris.Context, req Req_T, res Res_T) error {

	if ctx == nil {

		return ERR_NO_CONTEXT
	}

	req.ReceiveAccessToken(common.GetAccessToken(ctx))

	return nil
}

func UseAuthority[Req_T requestInput.IAuthorityBringAlong, Res_T any](ctx iris.Context, req Req_T, res Res_T) error {

	accessToken := common.GetAccessToken(ctx)

	if accessToken == nil {

		return nil
	}

	req.SetAuthority(accessToken.GetAuthData())
	return nil
}

func UseTenantMapping[Req_T requestInput.ITenantMappingInput, Res_T any](ctx iris.Context, req Req_T, res Res_T) error {

	accessToken := common.GetAccessToken(ctx)

	if accessToken == nil {

		return nil
	}

	req.SetTenantUUID(accessToken.GetTenantUUID())
	return nil
}
