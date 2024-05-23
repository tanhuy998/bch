package responsePresenter

import "app/domain/model"

type ModifyCandidateResponse struct {
	Message string           `json:"message"`
	Data    *model.Candidate `json:"data"`
}
