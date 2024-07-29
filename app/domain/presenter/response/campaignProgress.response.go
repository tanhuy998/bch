package responsePresenter

import "app/domain/valueObject"

type (
	CampaignProgressResponsePresenter struct {
		Data valueObject.CandidateSigningReport `json:"data"`
	}
)
