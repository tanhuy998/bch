package signingServicePort

import (
	"app/src/model"
	"app/src/repository"
)

type (
	IGetCampaignUnSignedCandidates interface {
		Serve(
			campaignUUID string,
			pivotObjID_str string,
			limit int,
			isPrevDir bool,
		) (*repository.PaginationPack[model.Candidate], error)
	}
)
