package repository

import (
	"app/domain/model"
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	CANDIDATE_SIGNING_INFO_COLLECTION = "candidateSigningInfos"
)

type (
	ICandidateSigningInfo interface {
		IMongoDBRepository
		GetOneByUUID(uuid uuid.UUID) (*model.CandidateSigningInfo, error)
		FindOneByCandidateUUID(candidateUUID uuid.UUID, ctx context.Context) (*model.CandidateSigningInfo, error)
		Create(signingInfo *model.CandidateSigningInfo, ctx context.Context) error
		Update(uuid *uuid.UUID, signingInfo *model.CandidateSigningInfo, ctx context.Context) error
	}

	CandidateSigningInfoRepository struct {
		AbstractMongoRepository
	}
)

func (this *CandidateSigningInfoRepository) GetDBClient() *mongo.Client {

	return this.collection.Database().Client()
}

func (this *CandidateSigningInfoRepository) GetCollection() *mongo.Collection {

	return this.collection
}

func (this *CandidateSigningInfoRepository) Init(db *mongo.Database) *CandidateSigningInfoRepository {

	this.AbstractMongoRepository.Init(db, CANDIDATE_SIGNING_INFO_COLLECTION)

	return this
}

func (this *CandidateSigningInfoRepository) GetOneByUUID(uuid uuid.UUID) (*model.CandidateSigningInfo, error) {

	return findOneDocument[model.CandidateSigningInfo](
		bson.D{
			{"uuid", uuid},
		},
		this.collection,
		context.TODO(),
	)
}

func (this *CandidateSigningInfoRepository) Update(
	uuid *uuid.UUID,
	signingInfo *model.CandidateSigningInfo,
	ctx context.Context,
) error {

	res, err := updateDocument[model.CandidateSigningInfo](
		uuid,
		signingInfo,
		this.collection,
		context.TODO(),
	)

	if err != nil {

		return err
	}

	return CheckUpdateOneResult(res)
}

func (this *CandidateSigningInfoRepository) FindOneByCandidateUUID(candidateUUID uuid.UUID, ctx context.Context) (*model.CandidateSigningInfo, error) {

	ret, err := findOneDocument[model.CandidateSigningInfo](
		bson.D{
			{"candidateUUID", candidateUUID},
		},
		this.collection,
		ctx,
	)

	if err == mongo.ErrNoDocuments {

		return nil, nil
	}

	if err != nil {

		return nil, err
	}

	return ret, nil
}

func (this *CandidateSigningInfoRepository) Create(signingInfo *model.CandidateSigningInfo, ctx context.Context) error {

	return createDocument[model.CandidateSigningInfo](signingInfo, this.collection, ctx)
}
