package signingService

import (
	adminServiceAdapter "app/adapter/adminService"
	"app/src/model"
	"app/src/repository"
	"context"

	"github.com/google/uuid"
)

type (
	IGetSingleCandidateSigningInfo interface {
		Serve(candidateUUID_str string) (*model.CandidateSigningInfo, error)
	}

	GetSingleCandidateSigningInfoService struct {
		CheckCandidateExistence adminServiceAdapter.ICheckCandidateExistence
		SigningInfoRepo         repository.ICandidateSigningInfo
	}
)

func (this *GetSingleCandidateSigningInfoService) Serve(
	candidateUUID_str string,
) (*model.CandidateSigningInfo, error) {

	candidateUUID, err := uuid.Parse(candidateUUID_str)

	if err != nil {

		return nil, err
	}

	// exist, err := this.CheckCandidateExistence.Serve(candidateUUID)

	// if err != nil {

	// 	return nil, err
	// }

	// if !exist {

	// 	return nil, ERR_CANDIDATE_NOT_FOUND
	// }

	data, err := this.SigningInfoRepo.FindOneByCandidateUUID(candidateUUID, context.TODO())

	if err != nil {

		return nil, err
	}

	return data, nil
}
