package presenter

import "app/src/model"

type GetSingleCandidateSigningInfoResponse struct {
	Message string                      `json:"message"`
	Data    *model.CandidateSigningInfo `json:"data"`
}
