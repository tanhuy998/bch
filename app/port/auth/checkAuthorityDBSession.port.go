package authServicePort

import (
	"app/internal/generalToken"
	"context"

	"github.com/google/uuid"
)

type (
	ICheckAuthorityDBSession interface {
		Serve(
			tenantUUID, userUUID uuid.UUID, sessionID generalToken.GeneralTokenID, ctx context.Context,
		) error
	}
)
