package valueObject

import "app/src/model"

type (
	CandidateLastCommit struct {
		Candidate  *model.Candidate
		LastCommit *model.CandidateSigningCommit
	}
)
