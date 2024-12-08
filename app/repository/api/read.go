package repositoryAPI

import (
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type (
	IRepositoryReadOperator[Model_T any] interface {
		FindMany(query bson.D, ctx context.Context) ([]*Model_T, error)
		FindOneByUUID(uuid uuid.UUID, ctx context.Context) (*Model_T, error)
		//Find(query bson.D, ctx context.Context) (*Model_T, error)
	}
)
