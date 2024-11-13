package genericUseCase

import (
	libCommon "app/internal/lib/common"
	"app/valueObject/requestInput"

	"app/model"
)

type (
	TenantDomainUseCase[Input_T requestInput.ITenantDomainInput, Output_T any] struct {
	}
)

func (TenantDomainUseCase[Input_T, Output_T]) GenerateTenantDomainModel(input Input_T) model.TenantDomainModel {

	auth := input.GetAuthority()

	return model.TenantDomainModel{
		TenantUUID: libCommon.Ternary(input.IsValidTenantUUID(), libCommon.PointerPrimitive(input.GetTenantUUID()), nil),
		CreatedBy:  libCommon.Ternary(auth == nil, nil, libCommon.PointerPrimitive(auth.GetUserUUID())),
	}
}
