package adminService

import (
	"app/domain/model"
	libCommon "app/lib/common"
	"app/repository"
	"context"
	"errors"
	"reflect"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

const (
	DEADLINE_DURATION = 5 * time.Second
)

type (
	IGetCampaignCandidateList interface {
		Serve(
			campaignUUID string,
			candiatePivot_id string,
			limit int,
			isPrevDir bool,
		) (*repository.PaginationPackWithHeader[model.Candidate, model.Campaign], error)
	}

	AdminGetCampaignCandidateListService struct {
		AbstractRepositoryInterface reflect.Type
		CampaignRepo                repository.ICampaignRepository
		CandidateRepo               repository.ICandidateRepository
	}
)

func (this *AdminGetCampaignCandidateListService) Serve(
	str_campaignUUID string, str_candiatePivot_id string, limit int, isPrevDir bool,
) (*repository.PaginationPackWithHeader[model.Candidate, model.Campaign], error) {

	campaignUUID, err := uuid.Parse(str_campaignUUID)

	if err != nil {

		return nil, err
	}

	candidateObjID, err := primitive.ObjectIDFromHex(str_candiatePivot_id)

	if err != nil {

		return nil, err
	}

	deadlineCtx, cancel := context.WithTimeout(context.Background(), DEADLINE_DURATION)
	defer cancel()

	session, err := initAggregateTransaction(this.CampaignRepo.GetDBClient())

	if err != nil {

		return nil, err
	}
	defer (*session).EndSession(deadlineCtx)

	return this.Query(
		session,
		campaignUUID,
		candidateObjID,
		int64(limit),
		isPrevDir,
		deadlineCtx,
	)
}

func (this *AdminGetCampaignCandidateListService) Query(
	session *mongo.Session,
	campaignUUID uuid.UUID,
	candidatePivot_id primitive.ObjectID,
	limit int64,
	isPrevDir bool,
	ctx context.Context,
) (*repository.PaginationPackWithHeader[model.Candidate, model.Campaign], error) {

	ctx = libCommon.Ternary(ctx == nil, context.TODO(), ctx)
	defer (*session).EndSession(ctx)

	pack, err := (*session).WithTransaction(
		ctx,
		func(sessionCtx mongo.SessionContext) (interface{}, error) {

			campaign, err := this.CampaignRepo.FindByUUID(campaignUUID, sessionCtx)

			if err != nil {

				return nil, err

			} else if campaign == nil {

				return nil, errors.New("campaign not found")
			}

			paginationPack, err := this.CandidateRepo.
				GetCandidaiteList(
					*campaign.UUID,
					candidatePivot_id,
					limit,
					isPrevDir,
					sessionCtx,
				)

			if err != nil {

				return nil, err
			}

			ret := &repository.PaginationPackWithHeader[model.Candidate, model.Campaign]{
				campaign, paginationPack,
			}

			return ret, nil
		},
		initAggregateTransactionOption(),
	)

	if err != nil {

		return nil, err
	}

	// if actualDataPack, ok := pack.(*repository.PaginationPack[model.Candidate]); ok {

	// 	return actualDataPack, nil
	// }

	if actualDataPack, ok := pack.(*repository.PaginationPackWithHeader[model.Candidate, model.Campaign]); ok {

		return actualDataPack, nil
	}

	return nil, errors.New("internal error")
}

func initAggregateTransactionOption() *options.TransactionOptions {

	writeConcern := writeconcern.Majority()
	readConcern := readconcern.Snapshot()
	return options.Transaction().SetWriteConcern(writeConcern).SetReadConcern(readConcern)
}

func initAggregateTransaction(dbClient *mongo.Client) (*mongo.Session, error) {

	session, err := dbClient.StartSession()

	if err != nil {

		return nil, err
	}

	return &session, nil
}
