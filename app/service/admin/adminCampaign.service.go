package adminService

import (
	"app/app/db"
	"app/app/model"
	"app/app/repository"
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

type CampaignAdminService struct {
	DBClient            *mongo.Client
	CandidateCollection *db.MongoCandidateCollection
	CampaignRepo        repository.ICampaignRepository
	CandidateRepo       repository.ICandidateRepository
}

func (this *CampaignAdminService) LaunchNewCampaign(model *model.Campaign) error {

	model.UUID = uuid.New()

	return this.CampaignRepo.Create(model, nil)
}

func (this *CampaignAdminService) ModifyExistingCampaign(uuid uuid.UUID, model *model.Campaign) error {

	model.UUID = uuid

	return this.CampaignRepo.Update(model, nil)
}

func (this *CampaignAdminService) DeleteCampaign(uuid uuid.UUID) error {

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
		delectCampain(uuid, this.CampaignRepo, this.CandidateCollection),
		transactionOpts,
	)

	return err
}

func delectCampain(
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
