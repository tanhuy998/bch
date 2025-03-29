package read

import (
	"app/internal/db/relation"
	"app/unitOfWork/aggregate/api/crud"
)

type (
	RelationReaderExecutor[Read_Entity_T, Local_Storage_Unit_Entity_T any] struct {
		QueryReaderExecutor[Read_Entity_T, Local_Storage_Unit_Entity_T]
	}
)

func (this *QueryReaderExecutor[Read_Entity_T, Local_Storage_Unit_Entity_T]) ByRelations(
	relation ...relation.IRelationForeignNavigator,
) crud.IAggregateReader[Read_Entity_T] {

	return this
}
