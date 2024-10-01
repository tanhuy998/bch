package responsePresenter

import "app/src/valueObject"

type (
	CampaignProgressResponsePresenter struct {
		Data valueObject.CandidateSigningReport `json:"data"`
	}
)
