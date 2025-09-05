package binding

import (
	libIris "app/internal/lib/iris"
	accessTokenClientPort "app/port/accessTokenClient"
	"app/valueObject/requestInput"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
)

type (
	Hook func(
		container *hero.Container, ctx iris.Context, req interface{}, res interface{},
	) error
)

func readAccessToken(
	ctx iris.Context, accessTokenClient accessTokenClientPort.IAccessTokenClient,
) error {

	at, err := accessTokenClient.Read(ctx)

	if err != nil {

		return err
	}

	libIris.SetAccessToken(ctx, at)

	return nil
}

func UseAuthority(
	container *hero.Container, ctx iris.Context, req interface{}, res interface{},
) error {

	switch input := (req).(type) {
	case requestInput.IAuthorityBringAlong:
		accessToken := libIris.GetAccessToken(ctx)

		if accessToken == nil {

			container.Handler(readAccessToken)(ctx)

			accessToken = libIris.GetAccessToken(ctx)
		}

		if accessToken == nil {

			return nil
		}

		input.SetAuthority(accessToken.GetAuthData())
		return nil
	default:
		return nil
	}
}

func UseTenantMapping(
	container *hero.Container, ctx iris.Context, req interface{}, res interface{},
) error {

	switch input := (req).(type) {
	case requestInput.ITenantMappingInput:

		accessToken := libIris.GetAccessToken(ctx)

		if accessToken == nil {

			container.Handler(readAccessToken)(ctx)

			accessToken = libIris.GetAccessToken(ctx)
		}

		if accessToken == nil {

			return nil
		}

		input.SetTenantUUID(accessToken.GetTenantUUID())
		return nil
	default:
		return nil
	}
}
