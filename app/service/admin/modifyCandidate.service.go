package adminService

import (
	"app/domain/model"
	"app/repository"

	"github.com/google/uuid"
)

type (
	IModifyCandidate interface {
		Execute(inputUUID string, model *model.Candidate) error
	}

	AdminModifyCandidate struct {
		CandidateRepoo repository.ICandidateRepository
	}
)

func (this *AdminModifyCandidate) Execute(inputUUID string, model *model.Candidate) error {

	uuid, err := uuid.Parse(inputUUID)

	if err != nil {

		return err
	}

	model.UUID = &uuid

	return this.CandidateRepoo.Update(model, nil)
}
