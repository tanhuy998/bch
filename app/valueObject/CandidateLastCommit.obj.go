package valueObject

import "app/model"

type (
	CandidateLastCommit struct {
		Candidate  *model.Candidate
		LastCommit *model.CandidateSigningCommit
	}
)
