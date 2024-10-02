package responsePresenter

import "app/valueObject"

type (
	CampaignProgressResponsePresenter struct {
		Data valueObject.CandidateSigningReport `json:"data"`
	}
)
