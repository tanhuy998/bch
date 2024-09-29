package responsePresenter

type (
	RefreshLoginData struct {
		AccessToken string `json:"accessToken"`
	}

	RefreshLoginResponse struct {
		Message string            `json:"message"`
		Data    *RefreshLoginData `json:"data,omitempty"`
	}
)
