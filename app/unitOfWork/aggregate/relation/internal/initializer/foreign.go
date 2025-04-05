package initializer

import (
	"app/internal/db/query"
	"app/unitOfWork/aggregate/relation/bridge/api"
)

type (
	Foreign[Initiator_T api.IForeign] struct {
		Initiator Initiator_T
	}
)

func (this Foreign[Initiator_T]) GetRelationInitFunc() query.JoinInitFunc {

	return this.Initiator.GetForeignInitFunc()
}

func (this Foreign[Initiator_T]) GetDBStorageUnitName() string {

	return this.Initiator.GetForeignDBStorageUnitName()
}
