package switchTenantDomain

import (
	"app/internal/common"
	libCommon "app/internal/lib/common"
	"app/model"
	accessTokenServicePort "app/port/accessToken"
	authSignatureTokenPort "app/port/authSignatureToken"
	generalTokenServicePort "app/port/generalToken"
	"app/repository"
	repositoryAPI "app/repository/api"
	"app/unitOfWork"
	"errors"
	"fmt"

	refreshTokenServicePort "app/port/refreshToken"
	"context"

	"github.com/google/uuid"
)

type (
	SwitchTenantService struct {
		unitOfWork.OperationLogger
		AuthSignatureTokenProvider authSignatureTokenPort.IAuthSignatureProvider
		UserSessionRepo            repository.IUserSession
	}
)

func (this *SwitchTenantService) Serve(
	tenantUUID uuid.UUID, generalToken generalTokenServicePort.IGeneralToken, ctx context.Context,
) (at accessTokenServicePort.IAccessToken, rt refreshTokenServicePort.IRefreshToken, err error) {

	if generalToken.Expire() {

		err = errors.Join(
			common.ERR_UNAUTHORIZED, fmt.Errorf("general token expires login again"),
		)
		return
	}

	// switch existingUserSession, e := this.UserSessionRepo.Find(
	// 	bson.D{
	// 		{"sessionID", generalToken.GetTokenID()},
	// 	},
	// 	ctx,
	// ); {
	// case e != nil:
	// 	err = e
	// 	return
	// case existingUserSession != nil:
	// 	err = errGeneralTokenMustBeRemovedFromClient
	// 	return
	// }

	query := this.UserSessionRepo.Filter(
		func(filter repositoryAPI.IFilterGenerator) {

			filter.Field("sesessionID").Equal(generalToken.GetTokenID())
		},
	)

	switch existingUserSession, e := query.FindOne(ctx); {
	case e != nil:
		err = e
		return
	case existingUserSession != nil:
		err = errGeneralTokenMustBeRemovedFromClient
		return
	}

	// err = this.UserSessionRepo.UpsertManyByFilter(
	// 	bson.D{
	// 		{"userUUID", generalToken.GetUserUUID()},
	// 		{"tenantUUID", tenantUUID},
	// 	},
	// 	bson.D{
	// 		{"userUUID", generalToken.GetUserUUID()},
	// 		{"tenantUUID", tenantUUID},
	// 		{"sessionID", generalToken.GetTokenID()},
	// 		{"expire", generalToken.GetExpiretime()},
	// 	},
	// 	ctx,
	// )

	err = this.UserSessionRepo.Filter(
		func(filter repositoryAPI.IFilterGenerator) {

			filter.Field("userUUID").Equal(generalToken.GetUserUUID())
			filter.Field("tenantUUID").Equal(tenantUUID)
		},
	).Upsert(
		model.UserSession{
			UserUUID:   libCommon.PointerPrimitive(generalToken.GetUserUUID()),
			TenantUUID: &tenantUUID,
			SessionID:  libCommon.PointerPrimitive(generalToken.GetTokenID()),
			Expire:     generalToken.GetExpiretime(),
		},
		ctx,
	)

	if err != nil {

		return
	}

	return this.AuthSignatureTokenProvider.Generate(tenantUUID, generalToken, ctx)
}
