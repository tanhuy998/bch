package adminService

import (
	"app/repository"
	"context"

	"github.com/google/uuid"
)

type (
	ICheckCandidateExistence interface {
		Serve(candidateUUID uuid.UUID) (bool, error)
	}

	CheckCandidateExistenceService struct {
		CandidateRepo repository.ICandidateRepository
	}
)

func (this *CheckCandidateExistenceService) Serve(candidateUUID uuid.UUID) (bool, error) {

	res, err := this.CandidateRepo.FindByUUID(candidateUUID, context.TODO())

	if err != nil {

		return false, err
	}

	if res == nil {

		return false, nil
	}

	return true, nil
}
