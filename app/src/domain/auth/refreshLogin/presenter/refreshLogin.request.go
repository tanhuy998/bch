package requestPresenter

import "github.com/kataras/iris/v12"

type (
	RefreshLoginRequest struct {
		Data struct {
			AccessToken string `json:"accessToken"`
		} `json:"data"`
		ctx iris.Context
	}
)

func (this *RefreshLoginRequest) ReceiveContext(ctx iris.Context) {

	this.ctx = ctx
}
func (this *RefreshLoginRequest) GetContext() iris.Context {

	return this.ctx
}
