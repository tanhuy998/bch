package requestInput

import "github.com/google/uuid"

type (
	TenantMappingInput struct {
		tenantUUID uuid.UUID
	}
)

func (this *TenantMappingInput) SetTenantUUID(tenantUUID uuid.UUID) {

	this.tenantUUID = tenantUUID
}

func (this *TenantMappingInput) IsValidTenantUUID() bool {

	return this.tenantUUID != uuid.Nil
}

func (this *TenantMappingInput) GetTenantUUID() uuid.UUID {

	return this.tenantUUID
}
