package signingService

import (
	"app/internal/common"
	"app/repository"
	"context"

	"github.com/google/uuid"
)

/*
 */
type (
	ICheckCandidateExistence interface {
		Serve(candidateUUID uuid.UUID) error
	}

	CheckCandidateExistenceService struct {
		CandidateRepo repository.ICandidateRepository
	}
)

func (this *CheckCandidateExistenceService) Serve(candidateUUID uuid.UUID) error {

	candidate, err := this.CandidateRepo.FindByUUID(candidateUUID, context.TODO())

	if err != nil {

		return err
	}

	if candidate == nil {

		return common.ERR_HTTP_NOT_FOUND
	}

	return nil
}
