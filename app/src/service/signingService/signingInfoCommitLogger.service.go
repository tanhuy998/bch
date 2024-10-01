package signingService

import (
	libCommon "app/src/internal/lib/common"
	"app/src/model"
	"app/src/repository"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/wI2L/jsondiff"
)

var (
	ERR_JSON_PATCH_NO_DIFF error = fmt.Errorf("no difference between json documents")
)

type (
	ISigningCommitLogger interface {
		Serve(
			signingInfoUUID *uuid.UUID,
			commitSigningInfo *model.CandidateSigningInfo,
			originalSigningInfo *model.CandidateSigningInfo,
		) error
		//CompareAndServe(commitSingingInfo *model.CandidateSigningInfo, candidate *model.Candidate) error
	}

	SigningCommmitLoggerService struct {
		CandidateRepository              repository.ICandidateRepository
		CandidateSigningCommitRepository repository.ICandidateSigningCommit
	}
)

func (this *SigningCommmitLoggerService) Serve(
	signingInfoUUID *uuid.UUID,
	commitSigningInfo *model.CandidateSigningInfo,
	originalSigningInfo *model.CandidateSigningInfo,
) error {

	return this.handle(signingInfoUUID, commitSigningInfo, originalSigningInfo)
}

// func (this *CandidateSigningCommmitLoggerService) CompareAndServe(
// 	commitSingingInfo *model.CandidateSigningInfo,
// 	originalSigningInfo *model.CandidateSigningInfo,
// ) error {

// 	return this.handle(commitSingingInfo, originalSigningInfo)
// }

func (this *SigningCommmitLoggerService) handle(
	singinInfoUUID *uuid.UUID,
	commitSingingInfo *model.CandidateSigningInfo,
	originalSigningInfo *model.CandidateSigningInfo,
) error {

	model, err := resolve(singinInfoUUID, commitSingingInfo, originalSigningInfo)

	if err != nil {

		return err
	}

	err = this.CandidateSigningCommitRepository.Create(model, nil)

	if err != nil {

		return err
	}

	return nil
}

func resolve(
	signingInfoUUID *uuid.UUID,
	commitSingingInfo *model.CandidateSigningInfo,
	existingSigningInfo *model.CandidateSigningInfo,
) (*model.CandidateSigningCommit, error) {

	jsonRawCommit, err := json.Marshal(commitSingingInfo)

	if err != nil {

		return nil, err
	}

	jsonRawOrginal, err := json.Marshal(existingSigningInfo)

	if err != nil {

		return nil, err
	}

	jsonPatch, err := jsondiff.CompareJSON(jsonRawOrginal, jsonRawCommit)

	if err != nil {

		return nil, err
	}

	if len(jsonPatch) == 0 {

		return nil, ERR_JSON_PATCH_NO_DIFF
	}

	ret := &model.CandidateSigningCommit{
		//CandidateUUID: &candidateUUID,
		SigningInfoUUID: signingInfoUUID,
		Time:            libCommon.PointerPrimitive(time.Now()),
		Operations:      jsonPatch.String(),
	}

	// err = json.Unmarshal(rawJsonPatch, &jsonPatchModel)

	// if err != nil {

	// 	return nil, err
	// }

	return ret, nil
}
