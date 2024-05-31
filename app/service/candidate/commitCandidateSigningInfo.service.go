package candidateService

import (
	"app/domain/model"
	"app/repository"

	"github.com/google/uuid"
)

type (
	ICommitCandidateSigningInfo interface {
		Serve(candidateUUID string, data *model.CandidateSigningInfo) error
	}

	CommitCandidateSigningInfoService struct {
		CandidateRepo repository.ICandidateRepository
	}
)

func (this *CommitCandidateSigningInfoService) Serve(
	candidateUUID_str string,
	data *model.CandidateSigningInfo,
) error {

	uuid, err := uuid.Parse(candidateUUID_str)

	if err != nil {

		return err
	}

	updateQuery := repository.CandidateSigninInfoUpdateQuery{
		SigningInfo: data,
	}

	this.CandidateRepo.UpdateSigningInfo(uuid, updateQuery, nil)

	return nil
}
