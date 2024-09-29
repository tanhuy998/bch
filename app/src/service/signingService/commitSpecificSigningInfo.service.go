package signingService

import (
	adminServiceAdapter "app/adapter/adminService"
	"app/domain/model"
	"app/internal/common"
	libCommon "app/lib/common"
	"app/repository"
	"context"
	"errors"

	"github.com/google/uuid"
)

type (
	ICommitSpecificSigningInfo interface {
		Serve(campaignUUID string, candidateUUID string, data *model.CandidateSigningInfo) error
	}

	CommitSpecificSigningInfoService struct {
		CandidateSigningInfoRepo       repository.ICandidateSigningInfo
		CheckCandidateExistService     ICheckCandidateExistence
		SigningCommmitLoggerService    ISigningCommitLogger
		CandidateRepo                  repository.ICandidateRepository
		AdminGetSingleCandidateAdapter adminServiceAdapter.IGetSingleCandidate
	}
)

var (
	ERR_CANDIDATE_NOT_FOUND = errors.New("CommitSpecificSigningInfo: candidate not found")
)

func (this *CommitSpecificSigningInfoService) Serve(campaignUUID_str string, candidateUUID_str string, data *model.CandidateSigningInfo) error {

	candidate, err := this.AdminGetSingleCandidateAdapter.Serve(candidateUUID_str)

	if err != nil {

		return err
	}

	if candidate == nil {

		return ERR_CANDIDATE_NOT_FOUND
	}

	campaignUUID, err := uuid.Parse(campaignUUID_str)

	if err != nil {

		return common.ERR_HTTP_NOT_FOUND
	}

	if *candidate.CampaignUUID != campaignUUID {

		return common.ERR_BAD_REQUEST
	}

	candidateUUID, err := uuid.Parse(candidateUUID_str)

	if err != nil {

		return errors.New("invalid input uuid")
	}

	existingSignInfo, err := this.CandidateSigningInfoRepo.FindOneByCandidateUUID(candidateUUID, context.TODO())

	if err != nil {

		return err
	}

	if existingSignInfo == nil {

		err := this.CheckCandidateExistService.Serve(candidateUUID)

		if err != nil {

			return err
		}
	}

	var signingInfoUUID *uuid.UUID
	/*
		signingInfoUUID has to be resolved for commit log
	*/
	if existingSignInfo != nil {

		signingInfoUUID = libCommon.PointerPrimitive(existingSignInfo.UUID)

	} else {

		signingInfoUUID = libCommon.PointerPrimitive(uuid.New())
	}

	err = this.writeCommitLog(signingInfoUUID, data, existingSignInfo)

	if err != nil {

		return err
	}

	if existingSignInfo != nil {

		err = this.CandidateSigningInfoRepo.Update(libCommon.PointerPrimitive(existingSignInfo.UUID), data, context.TODO())

	} else {
		/*
			signingInfoUUID already been resolved
		*/
		data.UUID = *signingInfoUUID
		data.CandidateUUID = candidateUUID
		data.CampaignUUID = campaignUUID

		err = this.CandidateSigningInfoRepo.Create(data, context.TODO())
	}

	if err != nil {

		return err
	}

	return nil
}

func (this *CommitSpecificSigningInfoService) writeCommitLog(
	signingInfoUUID *uuid.UUID,
	commitSigningInfo *model.CandidateSigningInfo,
	existingSignInfo *model.CandidateSigningInfo,
) error {

	return this.SigningCommmitLoggerService.Serve(signingInfoUUID, commitSigningInfo, existingSignInfo)
}
