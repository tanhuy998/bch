package authServiceAdapter

import (
	"app/domain/valueObject"
	"context"

	"github.com/google/uuid"
)

type (
	AuthData struct {
	}

	IFetchAuthData interface {
		Serve(userUUID uuid.UUID, ctx context.Context) (*valueObject.AuthData, error)
	}
)
