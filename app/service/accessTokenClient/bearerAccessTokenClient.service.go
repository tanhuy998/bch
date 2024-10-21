package accessTokenClientService

import (
	libError "app/internal/lib/error"
	accessTokenServicePort "app/port/accessToken"
	"context"
	"fmt"
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

func (this *BearerAccessTokenClientService) Read(ctx context.Context) (accessTokenServicePort.IAccessToken, error) {

	c, ok := ctx.(iris.Context)

	if !ok {

		return nil, libError.NewInternal(fmt.Errorf("context passed to BearerAccessTokenClientService is not type of iris.Context"))
	}

	header_value := c.GetHeader(HEADER_AUTH)

	if header_value == "" {

		return nil, nil
	}

	raw := strings.TrimPrefix(header_value, "Bearer ")

	return this.AccessTokenManipulator.Read(raw)
}
