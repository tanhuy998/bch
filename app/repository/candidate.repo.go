package repository

import (
	"app/app/model"

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

func (this *CandidateRepository) FindByUUID(uuid uuid.UUID) (*model.Campaign, error) {

	return findDocumentByUUID[model.Campaign](uuid, this.collection)
}

func (this *CandidateRepository) Get(page int) ([]*model.Campaign, error) {

	return getDocuments[model.Campaign](int64(page), this.collection)
}

func (this *CandidateRepository) Create(candidate *model.Candidate) error {

	return createDocument(candidate, this.collection)
}

func (this *CandidateRepository) Update(candidate *model.Candidate) error {

	return updateDocument(candidate.UUID, candidate, this.collection)
}

func (this *CandidateRepository) Delete(uuid uuid.UUID) error {

	return deleteDocument(uuid, this.collection)
}
