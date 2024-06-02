package responsePresenter

import "app/domain/model"

type GetSingleCandidateSigningInfoResponse struct {
	Message string
	Data    *model.CandidateSigningInfo
}
