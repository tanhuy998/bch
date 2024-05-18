package adminService

import (
	"app/app/repository"
	"context"

	"github.com/google/uuid"
)

type (
	IDeleteCandidate interface {
		Execute(string) error
	}

	AdminDeleteCandidateService struct {
		CandidateRepo repository.ICandidateRepository
	}
)

func (this *AdminDeleteCandidateService) Execute(inputUUID string) error {

	uuid, err := uuid.Parse(inputUUID)

	if err != nil {

		return err
	}

	return this.deleteCandidate(uuid)
}

func (this *AdminDeleteCandidateService) deleteCandidate(uuid uuid.UUID) error {

	return this.CandidateRepo.Delete(uuid, context.TODO())
}
