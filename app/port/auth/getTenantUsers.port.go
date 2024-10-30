package authServicePort

import (
	"app/model"
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	IGetTenantUsers interface {
		Serve(
			tenantUUID uuid.UUID, page uint64, size uint64, cursor *primitive.ObjectID, isPrev bool, ctx context.Context,
		) ([]model.User, error)
	}
)
