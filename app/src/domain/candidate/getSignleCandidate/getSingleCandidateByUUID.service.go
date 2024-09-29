package adminService

import (
	"app/domain/model"
	"app/repository"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type (
	IGetSingleCandidateByUUID interface {
		Serve(string) (*model.Candidate, error)
	}

	AdminGetSingleCandidateByUUIDService struct {
		CandidateRepo repository.ICandidateRepository
	}
)

func (this *AdminGetSingleCandidateByUUIDService) Serve(uuid_str string) (*model.Candidate, error) {

	uuid, err := uuid.Parse(uuid_str)

	if err != nil {

		return nil, err
	}

	return this.CandidateRepo.Find(
		bson.D{
			{"uuid", uuid},
		},
		nil,
	)
}
