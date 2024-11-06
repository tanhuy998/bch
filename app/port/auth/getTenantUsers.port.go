package authServicePort

import (
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	IGetTenantUsers[User_Entity_T any] interface {
		Serve(
			tenantUUID uuid.UUID, page uint64, size uint64, cursor *primitive.ObjectID, isPrev bool, ctx context.Context,
		) ([]User_Entity_T, error)
	}
)
