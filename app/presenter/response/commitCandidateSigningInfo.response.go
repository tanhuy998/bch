package responsePresenter

import "app/model"

type CommitCandidateSigningInfoResponse struct {
	Message     string           `json:"message"`
	UpdatedData *model.Candidate `json:"updatedData"`
}
