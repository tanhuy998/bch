package generalTokenIDServicePort

import (
	"app/internal/generalToken"

	"github.com/google/uuid"
)

type (
	IGeneralTokenIDProvider interface {
		Serve(userUUID uuid.UUID) (generalToken.GeneralTokenID, error)
	}
)
