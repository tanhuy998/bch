package accessTokenClientService

import (
	accessTokenServicePort "app/port/accessToken"
	"strings"

	"github.com/kataras/iris/v12"
)

const (
	HEADER_AUTH = "Authorization"
)

type (
	BearerAccessTokenClientService struct {
		AccessTokenManipulator accessTokenServicePort.IAccessTokenManipulator
	}
)

func (this *BearerAccessTokenClientService) Read(ctx iris.Context) (accessTokenServicePort.IAccessToken, error) {

	header_value := ctx.GetHeader(HEADER_AUTH)

	if header_value == "" {

		return nil, nil
	}

	raw := strings.TrimPrefix(header_value, "Bearer ")

	return this.AccessTokenManipulator.Read(raw)
}
