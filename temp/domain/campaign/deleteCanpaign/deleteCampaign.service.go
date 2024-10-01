package adminService

import (
	"app/infrastructure/db"
	"app/src/repository"
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

type (
	IDeleteCampaign interface {
		Execute(string) error
	}

	AdminDeleteCampaignService struct {
		DBClient            *mongo.Client
		CandidateCollection *db.MongoCandidateCollection
		CampaignRepo        repository.ICampaignRepository
		//CandidateRepo       repository.ICandidateRepository
	}
)

func (this *AdminDeleteCampaignService) Execute(inputUUID string) error {

	uuid, err := uuid.Parse(inputUUID)

	if err != nil {

		return err
	}

	return this.deleteCampain(uuid)
}

func (this *AdminDeleteCampaignService) deleteCampain(uuid uuid.UUID) error {

	session, err := this.DBClient.StartSession()

	if err != nil {

		return err
	}
	defer session.EndSession(context.TODO())

	writeConcern := writeconcern.Majority()
	readConcern := readconcern.Snapshot()
	transactionOpts := options.Transaction().SetWriteConcern(writeConcern).SetReadConcern(readConcern)

	_, err = session.WithTransaction(
		context.TODO(),
		delectOperation(uuid, this.CampaignRepo, this.CandidateCollection),
		transactionOpts,
	)

	return err
}

func delectOperation(
	uuid uuid.UUID,
	campaignRepo repository.ICampaignRepository,
	candidateCollection *db.MongoCandidateCollection,
) func(mongo.SessionContext) (interface{}, error) {

	return func(sessionCtx mongo.SessionContext) (interface{}, error) {

		_, err := candidateCollection.DeleteMany(sessionCtx, bson.D{})

		if err != nil {

			return nil, err
		}

		err = campaignRepo.Delete(uuid, sessionCtx)

		if err != nil {

			return nil, err
		}

		return nil, nil
	}
}
