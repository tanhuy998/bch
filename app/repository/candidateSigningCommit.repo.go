package repository

import (
	"app/domain/model"
	"context"
)

type (
	ICandidateSigningCommit interface {
		IMongoDBRepository
		// FindByCandidateUUID(candidateUUID uuid.UUID, ctx context.Context) ([]*model.JsonPatchRawValueOperation, error)
		// Find(query bson.D, ctx context.Context, projections ...bson.E) ([]*model.JsonPatchRawValueOperation, error)
		Create(jsonPatch *model.JsonPatchRawValueOperation, ctx context.Context) error
	}

	CandidateSingingCommit struct {
		AbstractMongoRepository
	}
)

// func (this *CandidateSingingCommit) Find(
// 	query bson.D,
// 	ctx context.Context,
// 	projections ...bson.E,
// ) ([]*model.JsonPatchRawValueOperation, error) {

// }

func (this *CandidateSingingCommit) Create(
	jsonPatch *model.JsonPatchRawValueOperation,
	ctx context.Context,
) error {

	ctx = initContext(ctx)

	_, err := this.collection.InsertOne(ctx, jsonPatch)

	if err != nil {

		return err
	}

	return nil
}

func initContext(ctx context.Context) context.Context {

	if ctx == nil {

		return context.TODO()
	}

	return ctx
}
