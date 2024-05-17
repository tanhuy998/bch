package repository

import (
	"app/app/model"
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	CANDIDATE_COLLECTION_NAME = "candidates"
)

type CandidateRepository struct {
	AbstractRepository
}

func (this *CandidateRepository) Init(db *mongo.Database) *CandidateRepository {

	this.AbstractRepository.Init(db, CANDIDATE_COLLECTION_NAME)

	return this
}

func (this *CandidateRepository) FindByUUID(uuid uuid.UUID, ctx context.Context) (*model.Campaign, error) {

	return findDocumentByUUID[model.Campaign](uuid, this.collection, ctx)
}

func (this *CandidateRepository) Get(page int, ctx context.Context) ([]*model.Campaign, error) {

	return getDocuments[model.Campaign](int64(page), this.collection, ctx)
}

func (this *CandidateRepository) Create(candidate *model.Candidate, ctx context.Context) error {

	return createDocument(candidate, this.collection, ctx)
}

func (this *CandidateRepository) Update(candidate *model.Candidate, ctx context.Context) error {

	return updateDocument(candidate.UUID, candidate, this.collection, ctx)
}

func (this *CandidateRepository) Delete(uuid uuid.UUID, ctx context.Context) error {

	return deleteDocument(uuid, this.collection, ctx)
}
