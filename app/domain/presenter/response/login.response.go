package responsePresenter

type (
	LoginResponse struct {
		Message string `json:"message"`
		Data    struct {
			AccessToken string `json:"accessToken,omitempty"`
		} `json:"data"`
	}
)
