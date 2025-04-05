package initializer

import (
	"app/internal/db/query"
	"app/unitOfWork/aggregate/relation/bridge/api"
)

type (
	Bridge[Initiator_T api.IBridge] struct {
		Initiator Initiator_T
	}
)

func (this Bridge[Initiator_T]) GetRelationInitFunc() query.JoinInitFunc {

	return this.Initiator.GetBridgeInitFunc()
}

func (this Bridge[Initiator_T]) GetDBStorageUnitName() string {

	return this.Initiator.GetBridgeDBStorageUnitName()
}
