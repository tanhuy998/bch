package requestPresenter

import (
	accessTokenServicePort "app/port/accessToken"
	"context"

	"github.com/google/uuid"
)

type (
	InputUser struct {
		//UUID          uuid.UUID `json:"uuid" bson:"uuid"`
		Name     string `json:"name" bson:"name" validate:"required"`
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
		//IsDeactivated bool   `json:"deactivated" bson:"deactivated"`
		//Info          UserInfo  `json:"userInfo" bson:"userInfo"`
	}

	CreateUserRequestPresenter struct {
		ctx        context.Context
		auth       accessTokenServicePort.IAccessTokenAuthData
		tenantUUID uuid.UUID
		Data       *InputUser `json:"data" validate:"required"`
	}
)

func (this *CreateUserRequestPresenter) ReceiveContext(ctx context.Context) {

	this.ctx = ctx
}

func (this *CreateUserRequestPresenter) GetContext() context.Context {

	return this.ctx
}

func (this *CreateUserRequestPresenter) GetAuthority() accessTokenServicePort.IAccessTokenAuthData {

	return this.auth
}

func (this *CreateUserRequestPresenter) SetAuthority(auth accessTokenServicePort.IAccessTokenAuthData) {

	this.auth = auth
}

func (this *CreateUserRequestPresenter) SetTenantUUID(tenantUUID uuid.UUID) {

	this.tenantUUID = tenantUUID
}

func (this *CreateUserRequestPresenter) IsValidTenantUUID() bool {

	return this.tenantUUID != uuid.Nil
}

func (this *CreateUserRequestPresenter) GetTenantUUID() uuid.UUID {

	return this.tenantUUID
}
