package generalTokenIDService

import (
	"app/internal/generalToken"

	"github.com/google/uuid"
)

type (
	GeneralTokenIDProvider struct {
	}
)

func (this *GeneralTokenIDProvider) Serve(userUUID uuid.UUID) (id generalToken.GeneralTokenID, err error) {

	return generalToken.New(userUUID)
}
