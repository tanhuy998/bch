package requestPresenter

import (
	"app/model"
	accessTokenServicePort "app/port/accessToken"

	"github.com/google/uuid"
	"github.com/kataras/iris/v12"
)

type (
	CreateAssignmentGroupRequest struct {
		AssignmentUUID *uuid.UUID             `param:"assignmnetUUID" validate:"required"`
		Data           *model.AssignmentGroup `json:"data" validate:"required"`
		ctx            iris.Context
		auth           accessTokenServicePort.IAccessTokenAuthData
	}
)

func (this *CreateAssignmentGroupRequest) ReceiveContext(ctx iris.Context) {

	this.ctx = ctx
}

func (this *CreateAssignmentGroupRequest) GetContext() iris.Context {

	return this.ctx
}

func (this *CreateAssignmentGroupRequest) GetAuthority() accessTokenServicePort.IAccessTokenAuthData {

	return this.auth
}

func (this *CreateAssignmentGroupRequest) SetAuthority(auth accessTokenServicePort.IAccessTokenAuthData) {

	this.auth = auth
}
