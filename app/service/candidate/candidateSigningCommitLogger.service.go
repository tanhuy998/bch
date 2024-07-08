package candidateService

import (
	"app/domain/model"
	"app/repository"
	"encoding/json"

	jsonpatch "github.com/evanphx/json-patch/v5"
	"github.com/google/uuid"
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
) (*model.JsonPatchRawValueOperation, error) {

	jsonRawCommit, err := json.Marshal(commitSingingInfo)

	if err != nil {

		return nil, err
	}

	jsonRawOrginal, err := json.Marshal(existingSigningInfo)

	if err != nil {

		return nil, err
	}

	rawJsonPatch, err := jsonpatch.CreateMergePatch(jsonRawOrginal, jsonRawCommit)

	if err != nil {

		return nil, err
	}

	var jsonPatchModel *model.JsonPatchRawValueOperation

	err = json.Unmarshal(rawJsonPatch, &jsonPatchModel)

	if err != nil {

		return nil, err
	}

	return jsonPatchModel, nil
}
