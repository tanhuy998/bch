package checkAuthorityDomain

import (
	"app/internal/common"
	"app/internal/generalToken"
	"app/repository"
	repositoryAPI "app/repository/api"
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/kataras/iris/v12/x/errors"
)

type (
	CheckAuthorityDBSessionService struct {
		UserSessionRepo repository.IUserSession
	}
)

func (this *CheckAuthorityDBSessionService) Serve(
	tenantUUID, userUUID uuid.UUID, sessionID generalToken.GeneralTokenID, ctx context.Context,
) error {

	// res, err := this.UserSessionRepo.Find(
	// 	bson.D{
	// 		{"userUUID", userUUID},
	// 		{"sessionID", sessionID},
	// 	},
	// 	ctx,
	// )

	res, err := this.UserSessionRepo.Filter(
		func(filter repositoryAPI.IFilterGenerator) {

			filter.Field("userUUID").Equal(userUUID)
			filter.Field("sessionID").Equal(sessionID)
		},
	).FindOne(ctx)

	if err != nil {

		return err
	}

	if res == nil {

		return errors.Join(common.ERR_UNAUTHORIZED, fmt.Errorf("CheckAuthorityService user session deactivated"))
	}

	return nil
}
