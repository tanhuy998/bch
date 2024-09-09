package middlewareHelper

import (
	accessTokenServicePort "app/adapter/accessToken"
	"app/internal"
	"errors"

	"github.com/kataras/iris/v12"
)

var (
	ERR_NO_CONTEXT = errors.New("bindPresenters helper: no context")
)

type (
	PresenterInitializer[RequestPresenter_T, ResponsePresenter_T any] func(req *RequestPresenter_T, res *ResponsePresenter_T) error
	RequestPresenterInitializer[RequestPresenter_T any]               func(req *RequestPresenter_T) error

	IAccessTokenBringAlong interface {
		IContextBringAlong
		ReceiveAccessToken(at accessTokenServicePort.IAccessToken)
		GetAccessToken() accessTokenServicePort.IAccessToken
	}

	IContextBringAlong interface {
		ReceiveContext(ctx iris.Context)
		GetContext() iris.Context
	}
)

func BringAlongAccessToken[Req_T IAccessTokenBringAlong, Res_T any](req Req_T, res Res_T) error {

	ctx := req.GetContext()

	if ctx == nil {

		return ERR_NO_CONTEXT
	}

	unknown := ctx.Values().Get(internal.CTX_ACCESS_TOKEN_KEY)

	if accessToken, ok := unknown.(accessTokenServicePort.IAccessToken); ok {

		req.ReceiveAccessToken(accessToken)
	}

	return nil
}
