package removeDBUserSessionDomain

import (
	libCommon "app/internal/lib/common"
	"app/model"
	"app/repository"
	"context"

	"github.com/google/uuid"
)

type (
	RemoveDBUserSessionService struct {
		UserSessionRepo repository.IUserSession
	}
)

func (this *RemoveDBUserSessionService) Serve(userUUID uuid.UUID, ctx context.Context) error {

	err := this.UserSessionRepo.DeleteMany(
		&model.UserSession{
			UserUUID: libCommon.PointerPrimitive(userUUID),
		},
		ctx,
	)

	if err != nil {

		return err
	}

	return nil
}
