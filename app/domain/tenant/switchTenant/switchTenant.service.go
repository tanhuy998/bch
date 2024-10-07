package switchTenantDomain

import (
	accessTokenServicePort "app/port/accessToken"
	authSignatureTokenPort "app/port/authSignatureToken"
	"app/port/generalTokenServicePort"
	refreshTokenServicePort "app/port/refreshToken"
	"context"

	"github.com/google/uuid"
)

type (
	SwitchTenantService struct {
		AuthSignatureTokenProvider authSignatureTokenPort.IAuthSignatureProvider
	}
)

func (this *SwitchTenantService) Serve(
	tenantUUID uuid.UUID, generalToken generalTokenServicePort.IGeneralToken, ctx context.Context,
) (at accessTokenServicePort.IAccessToken, rt refreshTokenServicePort.IRefreshToken, err error) {

	return this.AuthSignatureTokenProvider.Generate(tenantUUID, generalToken, ctx)
}
