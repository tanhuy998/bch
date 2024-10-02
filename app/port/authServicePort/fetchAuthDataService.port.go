package authServicePort

import (
	"app/valueObject"
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
