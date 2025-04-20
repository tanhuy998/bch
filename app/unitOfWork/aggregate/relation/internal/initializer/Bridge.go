package initializer

import (
	"app/internal/db/query"
	"app/unitOfWork/aggregate/relation/binary/api"
)

type (
	Bridge[Initializer_T api.IBridge] struct {
		Initializer Initializer_T
	}
)

func (this Bridge[Initializer_T]) GetRelationInitFunc() query.JoinInitFunc {

	return this.Initializer.GetBridgeInitFunc()
}

func (this Bridge[Initializer_T]) GetDBStorageUnitName() string {

	return this.Initializer.GetBridgeDBStorageUnitName()
}
