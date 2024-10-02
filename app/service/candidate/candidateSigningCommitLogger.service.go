package candidateService

import (
	libCommon "app/internal/lib/common"
	"app/model"
	"app/repository"
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
	ICandidateSigningCommitLogger interface {
		Serve(candidateUUID string, commitSigningInfo *model.CandidateSigningInfo) error
		CompareAndServe(commitSingingInfo *model.CandidateSigningInfo, candidate *model.Candidate) error
	}

	CandidateSigningCommmitLoggerService struct {
		CandidateRepository              repository.ICandidateRepository
		CandidateSigningCommitRepository repository.ICandidateSigningCommit
	}
)

func (this *CandidateSigningCommmitLoggerService) Serve(
	candidateUUID_str string,
	commitSigningInfo *model.CandidateSigningInfo,
) error {

	uuid, err := uuid.Parse(candidateUUID_str)

	if err != nil {

		return err
	}

	candidate, err := this.CandidateRepository.FindByUUID(uuid, nil)

	if err != nil {

		return err
	}

	return this.handle(commitSigningInfo, candidate)
}

func (this *CandidateSigningCommmitLoggerService) CompareAndServe(
	commitSingingInfo *model.CandidateSigningInfo,
	candidate *model.Candidate,
) error {

	return this.handle(commitSingingInfo, candidate)
}

func (this *CandidateSigningCommmitLoggerService) handle(
	commitSingingInfo *model.CandidateSigningInfo,
	candidate *model.Candidate,
) error {

	model, err := resolve(*candidate.UUID, commitSingingInfo, candidate.SigningInfo)

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
	candidateUUID uuid.UUID,
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
		CandidateUUID: &candidateUUID,
		Time:          libCommon.PointerPrimitive(time.Now()),
		Operations:    jsonPatch.String(),
	}

	// err = json.Unmarshal(rawJsonPatch, &jsonPatchModel)

	// if err != nil {

	// 	return nil, err
	// }

	return ret, nil
}
