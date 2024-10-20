package requestPresenter

import (
	"app/valueObject/requestInput"
)

type (
	RefreshLoginRequest struct {
		requestInput.ContextInput
		Data struct {
			AccessToken string `json:"accessToken"`
		} `json:"data"`
	}
)
