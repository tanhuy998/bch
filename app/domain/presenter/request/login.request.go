package requestPresenter

import (
	refreshTokenServicePort "app/adapter/refreshToken"

	"github.com/kataras/iris/v12"
)

type (
	LoginInputUser struct {
		//UUID          uuid.UUID `json:"uuid" bson:"uuid"`
		//Name     string `json:"name" bson:"name" validate:"required"`
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
		//IsDeactivated bool   `json:"deactivated" bson:"deactivated"`
		//Info          UserInfo  `json:"userInfo" bson:"userInfo"`
	}

	LoginRequest struct {
		Data         LoginInputUser `json:"data"`
		ctx          iris.Context
		refreshToken refreshTokenServicePort.IRefreshToken
	}
)

func (this *LoginRequest) ReceiveContext(ctx iris.Context) {

	this.ctx = ctx
}
func (this *LoginRequest) GetContext() iris.Context {

	return this.ctx
}

func (this *LoginRequest) ReceiveRefreshToken(token refreshTokenServicePort.IRefreshToken) {

	this.refreshToken = token
}

func (this *LoginRequest) GetRefreshToken() refreshTokenServicePort.IRefreshToken {

	return this.refreshToken
}
