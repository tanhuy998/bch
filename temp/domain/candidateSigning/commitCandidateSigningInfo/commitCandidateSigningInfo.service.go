package candidateService

import (
	"app/src/model"
	"app/src/repository"

	"github.com/google/uuid"
)

type (
	ICommitCandidateSigningInfo interface {
		Serve(candidateUUID_str string, campaignUUID_str string, data *model.CandidateSigningInfo) error
	}

	CommitCandidateSigningInfoService struct {
		CandidateRepo repository.ICandidateRepository
	}
)

func (this *CommitCandidateSigningInfoService) Serve(
	candidateUUID_str string,
	campaignUUID_str string,
	data *model.CandidateSigningInfo,
) error {

	candidateUUID, err := uuid.Parse(candidateUUID_str)

	if err != nil {

		return err
	}

	campaignUUID, err := uuid.Parse(campaignUUID_str)

	if err != nil {

		return err
	}

	updateQuery := repository.CandidateSigninInfoUpdateQuery{
		SigningInfo: data,
	}

	err = this.CandidateRepo.UpdateSigningInfo(candidateUUID, campaignUUID, &updateQuery, nil)

	if err != nil {

		return err
	}

	return nil
}
