package adminService

import (
	"app/app/model"
	"app/app/repository"

	"github.com/google/uuid"
)

type CandidateAdminService struct {
	candidateRepo repository.ICandidateRepository
}

func (this *CandidateAdminService) ModifyExistingCandidate(uuid uuid.UUID, model *model.Candidate) error {

	model.UUID = uuid

	return this.candidateRepo.Update(model, nil)
}
func (this *CandidateAdminService) DeleteCandidate(uuid uuid.UUID) error {

	return this.candidateRepo.Delete(uuid, nil)
}
