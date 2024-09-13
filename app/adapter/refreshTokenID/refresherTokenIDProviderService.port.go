package refreshtokenidServicePort

import "github.com/google/uuid"

type (
	IRefreshTokenIDProvider interface {
		Generate(userUUID uuid.UUID) (string, error)
	}
)
