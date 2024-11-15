package switchTenantDomain

import (
	"app/internal/common"
	accessTokenServicePort "app/port/accessToken"
	authSignatureTokenPort "app/port/authSignatureToken"
	generalTokenServicePort "app/port/generalToken"
	"app/repository"
	"errors"
	"fmt"

	refreshTokenServicePort "app/port/refreshToken"
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type (
	SwitchTenantService struct {
		AuthSignatureTokenProvider authSignatureTokenPort.IAuthSignatureProvider
		UserSessionRepo            repository.IUserSession
	}
)

func (this *SwitchTenantService) Serve(
	tenantUUID uuid.UUID, generalToken generalTokenServicePort.IGeneralToken, ctx context.Context,
) (at accessTokenServicePort.IAccessToken, rt refreshTokenServicePort.IRefreshToken, err error) {

	if generalToken.Expire() {

		err = errors.Join(
			common.ERR_UNAUTHORIZED, fmt.Errorf("session expires login again"),
		)
		return
	}

	err = this.UserSessionRepo.UpsertManyByFilter(
		bson.D{
			{"userUUID", generalToken.GetUserUUID()},
			{"tenantUUID", tenantUUID},
		},
		bson.D{
			{"userUUID", generalToken.GetUserUUID()},
			{"tenantUUID", tenantUUID},
			{"sessionID", generalToken.GetTokenID()},
			{"expire", generalToken.GetExpiretime()},
		},
		ctx,
	)

	if err != nil {

		return
	}

	return this.AuthSignatureTokenProvider.Generate(tenantUUID, generalToken, ctx)
}
