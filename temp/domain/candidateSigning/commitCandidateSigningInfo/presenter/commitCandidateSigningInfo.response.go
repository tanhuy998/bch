package presenter

import "app/src/model"

type CommitCandidateSigningInfoResponse struct {
	Message     string           `json:"message"`
	UpdatedData *model.Candidate `json:"updatedData"`
}
