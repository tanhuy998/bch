package requestPresenter

import "context"

type (
	RefreshLoginRequest struct {
		Data struct {
			AccessToken string `json:"accessToken"`
		} `json:"data"`
		ctx context.Context
	}
)

func (this *RefreshLoginRequest) ReceiveContext(ctx context.Context) {

	this.ctx = ctx
}
func (this *RefreshLoginRequest) GetContext() context.Context {

	return this.ctx
}
