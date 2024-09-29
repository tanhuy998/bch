package responsePresenter

import "app/domain/model"

type GetSingleCandidateSigningInfoResponse struct {
	Message string                      `json:"message"`
	Data    *model.CandidateSigningInfo `json:"data"`
}
