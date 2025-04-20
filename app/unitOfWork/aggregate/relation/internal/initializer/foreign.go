package initializer

import (
	"app/internal/db/query"
	"app/unitOfWork/aggregate/relation/binary/api"
)

type (
	Foreign[Initializer_T api.IForeign] struct {
		Initializer Initializer_T
	}
)

func (this Foreign[Initializer_T]) GetRelationInitFunc() query.JoinInitFunc {

	return this.Initializer.GetForeignInitFunc()
}

func (this Foreign[Initializer_T]) GetDBStorageUnitName() string {

	return this.Initializer.GetForeignDBStorageUnitName()
}
