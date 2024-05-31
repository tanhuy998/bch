package repository

import (
	"app/domain/model"
	libCommon "app/lib/common"
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	CANDIDATE_COLLECTION_NAME  = "candidates"
	CANDIDATE_SIGNING_INFO_KEY = "signingInfo"
)

type (
	CandidateRepository struct {
		AbstractMongoRepository
	}
)

func (this *CandidateRepository) Init(db *mongo.Database) *CandidateRepository {

	this.AbstractMongoRepository.Init(db, CANDIDATE_COLLECTION_NAME)

	return this
}

/*
# IMPLEMENT AbstractMongoRepository
*/

func (this *CandidateRepository) GetCollection() *mongo.Collection {

	return this.collection
}

/*
# END IMPLEMENT AbstractMongoRepository
*/

/*
# IMPLEMENT ICandidateRepository
*/

func (this *CandidateRepository) GetCandidaiteList(
	campaignUUID uuid.UUID,
	pivot_id primitive.ObjectID,
	pageLimit int64,
	isPrevDir bool,
	ctx context.Context,
) (*PaginationPack[model.Candidate], error) {

	return getDocumentsPageByID[model.Candidate](
		pivot_id,
		pageLimit,
		isPrevDir,
		&bson.D{{"campaignUUID", 0}, {"signingInfo", 0}},
		this.collection,
		ctx,
		bson.E{"campaignUUID", campaignUUID},
	)
}

func (this *CandidateRepository) GetDBClient() *mongo.Client {

	return this.collection.Database().Client()
}

func (this *CandidateRepository) FindByUUID(uuid uuid.UUID, ctx context.Context) (*model.Candidate, error) {

	return findDocumentByUUID[model.Candidate](uuid, this.collection, ctx, bson.E{CANDIDATE_SIGNING_INFO_KEY, 0})
}

func (this *CandidateRepository) Get(page int, ctx context.Context) ([]*model.Candidate, error) {

	return getDocuments[model.Candidate](int64(page), this.collection, ctx)
}

func (this *CandidateRepository) Create(candidate *model.Candidate, ctx context.Context) error {

	return createDocument(candidate, this.collection, ctx)
}

func (this *CandidateRepository) Update(candidate *model.Candidate, ctx context.Context) error {

	return updateDocument(candidate.UUID, candidate, this.collection, ctx)
}

func (this *CandidateRepository) UpdateSigningInfo(
	candidateUUID uuid.UUID,
	query *CandidateSigninInfoUpdateQuery,
	ctx context.Context,
) error {

	return updateDocument(
		libCommon.PointerPrimitive(candidateUUID),
		query,
		this.collection,
		ctx,
	)
}

func (this *CandidateRepository) Delete(uuid uuid.UUID, ctx context.Context) error {

	return deleteDocument(uuid, this.collection, ctx)
}

/*
# END IMPLEMENT ICandidateRepository
*/
