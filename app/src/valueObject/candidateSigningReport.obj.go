package valueObject

type (
	CandidateSigningReport struct {
		TotalCount  int64 `json:"totalCount" bson:"totalCount"`
		SignedCount int64 `json:"signedCount" bson:"signedCount"`
	}

	CampaignCandidateCount struct {
		TotalCount int64 `json:"totalCount" bson:"totalCount"`
	}

	CampaignSignedCandidateCount struct {
		SignedCount int64 `json:'signedCount" bson:"signedCount"`
	}
)
