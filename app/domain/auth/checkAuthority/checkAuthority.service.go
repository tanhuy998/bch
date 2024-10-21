package checkAuthorityDomain

import (
	"app/internal/common"
	"app/internal/generalToken"
	"app/repository"
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/kataras/iris/v12/x/errors"
	"go.mongodb.org/mongo-driver/bson"
)

type (
	CheckAuthorityService struct {
		UserSessionRepo repository.IUserSession
	}
)

func (this *CheckAuthorityService) Serve(
	tenantUUID, userUUID uuid.UUID, sessionID generalToken.GeneralTokenID, ctx context.Context,
) error {

	res, err := this.UserSessionRepo.Find(
		bson.D{
			{"userUUID", userUUID},
			{"sessionID", sessionID},
		},
		ctx,
	)

	if err != nil {

		return err
	}

	if res == nil {

		return errors.Join(common.ERR_UNAUTHORIZED, fmt.Errorf("CheckAuthorityService user session deactivated"))
	}

	return nil
}
