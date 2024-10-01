package adminService

import (
	"app/src/model"
	"app/src/repository"

	"github.com/google/uuid"
)

type (
	IModifyExistingCandidate interface {
		Serve(inputUUID string, model *model.Candidate) error
	}

	AdminModifyExistingCandidate struct {
		CandidateRepoo repository.ICandidateRepository
	}
)

func (this *AdminModifyExistingCandidate) Serve(inputUUID string, model *model.Candidate) error {

	uuid, err := uuid.Parse(inputUUID)

	if err != nil {

		return err
	}

	model.UUID = &uuid

	return this.CandidateRepoo.Update(model, nil)
}
