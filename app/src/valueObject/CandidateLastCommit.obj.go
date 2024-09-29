package valueObject

import "app/domain/model"

type (
	CandidateLastCommit struct {
		Candidate  *model.Candidate
		LastCommit *model.CandidateSigningCommit
	}
)
