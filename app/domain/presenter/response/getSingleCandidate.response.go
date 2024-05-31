package responsePresenter

import "app/domain/model"

type GetSingleCandidateResponse struct {
	Message string           `json:"message"`
	Data    *model.Candidate `json:"data"`
}
