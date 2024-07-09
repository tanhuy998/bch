package repository

import (
	"app/domain/model"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type (
	ICandidateSigningCommit interface {
		// FindByCandidateUUID(candidateUUID uuid.UUID, ctx context.Context) ([]*model.JsonPatchRawValueOperation, error)
		// Find(query bson.D, ctx context.Context, projections ...bson.E) ([]*model.JsonPatchRawValueOperation, error)
		Create(jsonPatch *model.CandidateSigningCommit, ctx context.Context) error
	}

	CandidateSingingCommitRepository struct {
		AbstractMongoRepository
	}
)

// func (this *CandidateSingingCommit) Find(
// 	query bson.D,
// 	ctx context.Context,
// 	projections ...bson.E,
// ) ([]*model.JsonPatchRawValueOperation, error) {

// }

func (this *CandidateSingingCommitRepository) Init(db *mongo.Database) *CandidateSingingCommitRepository {

	this.AbstractMongoRepository.Init(db, "candidateSigningCommits")

	return this
}

func (this *CandidateSingingCommitRepository) Create(
	jsonPatch *model.CandidateSigningCommit,
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
