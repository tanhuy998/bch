package relations

import (
	"app/internal/db/query"
	"app/internal/db/relation"
	"app/model"
	"app/unitOfWork/aggregate/relation/triat/leftJoin"

	"github.com/google/uuid"
)

type (
	Tenant_TenanAgent_Relation struct {
		leftJoin.OneToManyWith[model.TenantAgent]
		userUUID uuid.UUID
	}
)

func (this Tenant_TenanAgent_Relation) InitializeForeign(join query.IJoinField) {

	join.On("uuid", "tenantUUID").As("tenantAgent").
		Filter(
			func(filter query.IFilterExpression) {
				filter.Field("userUUID").Equal(this.userUUID)
			},
		)
}

func (this Tenant_TenanAgent_Relation) GetRelationInitFunc() query.JoinInitFunc {

	return nil
}

func (this *Tenant_TenanAgent_Relation) DispatchDefault(
	userUUID uuid.UUID,
) relation.IDBRelationInitiator {

	clone := *this

	clone.userUUID = userUUID

	return clone
}
