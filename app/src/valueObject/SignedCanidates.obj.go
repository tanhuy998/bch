package valueObject

import "app/src/model"

type (
	SignedCandidates struct {
		List  []*model.Candidate `json:"list" bson:"list"`
		Count int64              `json:"count" bson:"count"`
	}
)
