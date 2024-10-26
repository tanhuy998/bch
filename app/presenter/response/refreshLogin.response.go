package responsePresenter

import "app/model"

type (
	RefreshLoginData struct {
		AccessToken string      `json:"accessToken"`
		User        *model.User `json:"user,omitempty"`
	}

	RefreshLoginResponse struct {
		Message string            `json:"message"`
		Data    *RefreshLoginData `json:"data,omitempty"`
	}
)
