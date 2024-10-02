package accessTokenClientService

import (
	"strings"

	"github.com/kataras/iris/v12"
)

const (
	HEADER_AUTH = "Authorization"
)

type (
	BearerAccessTokenClientService struct {
	}
)

func (this *BearerAccessTokenClientService) Read(ctx iris.Context) string {

	header_value := ctx.GetHeader(HEADER_AUTH)

	if header_value == "" {

		return ""
	}

	return strings.TrimPrefix(header_value, "Bearer ")
}
