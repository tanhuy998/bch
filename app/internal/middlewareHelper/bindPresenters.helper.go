package middlewareHelper

import (
	accessTokenServicePort "app/adapter/accessToken"
	"app/internal/common"
	"errors"

	"github.com/kataras/iris/v12"
)

var (
	ERR_NO_CONTEXT = errors.New("bindPresenters helper: no context")
)

type (
	PresenterInitializer[RequestPresenter_T, ResponsePresenter_T any] func(ctx iris.Context, req *RequestPresenter_T, res *ResponsePresenter_T) error
	RequestPresenterInitializer[RequestPresenter_T any]               func(req *RequestPresenter_T) error

	IAccessTokenBringAlong interface {
		//IContextBringAlong
		ReceiveAccessToken(at accessTokenServicePort.IAccessToken)
		GetAccessToken() accessTokenServicePort.IAccessToken
	}

	IContextBringAlong interface {
		ReceiveContext(ctx iris.Context)
		GetContext() iris.Context
	}

	IAuthorityBringAlong interface {
		GetAuthority() accessTokenServicePort.IAccessTokenAuthData
		SetAuthority(auth accessTokenServicePort.IAccessTokenAuthData)
	}
)

func UseAccessToken[Req_T IAccessTokenBringAlong, Res_T any](ctx iris.Context, req Req_T, res Res_T) error {

	if ctx == nil {

		return ERR_NO_CONTEXT
	}

	req.ReceiveAccessToken(common.GetAccessToken(ctx))

	return nil
}

func UseAuthority[Req_T IAuthorityBringAlong, Res_T any](ctx iris.Context, req Req_T, res Res_T) error {

	accessToken := common.GetAccessToken(ctx)

	if accessToken == nil {

		return nil
	}

	req.SetAuthority(accessToken.GetAuthData())
	return nil
}
