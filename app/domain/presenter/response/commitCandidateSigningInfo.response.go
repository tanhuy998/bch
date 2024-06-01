package responsePresenter

import "app/domain/model"

type CommitCandidateSigningInfoResponse struct {
	Message     string           `json:"message"`
	UpdatedData *model.Candidate `json:"updatedData"`
}
