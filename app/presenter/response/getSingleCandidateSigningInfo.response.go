package responsePresenter

import "app/model"

type GetSingleCandidateSigningInfoResponse struct {
	Message string                      `json:"message"`
	Data    *model.CandidateSigningInfo `json:"data"`
}
