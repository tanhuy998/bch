package signingService

import (
	"app/domain/model"
	libCommon "app/lib/common"
	"app/repository"
	"context"
	"errors"

	"github.com/google/uuid"
)

type (
	ICommitSpecificSigningInfo interface {
		Serve(candidateUUID string, data *model.CandidateSigningInfo) error
	}

	CommitSpecificSigningInfoService struct {
		CandidateSigningInfoRepo    repository.ICandidateSigningInfo
		CheckCandidateExistService  ICheckCandidateExistence
		SigningCommmitLoggerService ISigningCommitLogger
		CandidateRepo               repository.ICandidateRepository
	}
)

func (this *CommitSpecificSigningInfoService) Serve(candidateUUID_str string, data *model.CandidateSigningInfo) error {

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
