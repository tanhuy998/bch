package responsePresenter

import "app/src/model"

type GetSingleCandidateResponse struct {
	Message string           `json:"message"`
	Data    *model.Candidate `json:"data"`
}
