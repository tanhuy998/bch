package responsePresenter

import "app/model"

type GetSingleCandidateResponse struct {
	Message string           `json:"message"`
	Data    *model.Candidate `json:"data"`
}
