package signingServicePort

import (
	"app/model"
	"app/repository"
)

type (
	IGetCampaignSignedCandidates interface {
		Serve(
			campaignUUID string,
			pivotObjID_str string,
			limit int,
			isPrevDir bool,
		) (*repository.PaginationPack[model.Candidate], error)
	}
)
