package bridgeRelation

import (
	"app/internal/db/query"
	"app/internal/db/relation"
	"fmt"
)

type (
	BinaryBridgeRelation[
		Bridge_Initiator_T relation.IDBRelationInitiator,
		Foreign_Initiator_T relation.IDBRelationInitiator,
	] struct {
		bridge  Bridge_Initiator_T
		foreign Foreign_Initiator_T
	}
)

/*
Implemnt
*/
func (this *BinaryBridgeRelation[Bridge_Initiator_T, Foreign_Initiator_T]) GetRelationInitFunc() query.JoinInitFunc {

	defer func() {

		if r := recover(); r != nil {

			panic(fmt.Sprintf("BinaryBridgeRelation panic, %s", r))
		}
	}()

	return this.bridge.GetRelationInitFunc()
}

func (this *BinaryBridgeRelation[Bridge_Initiator_T, Foreign_Initiator_T]) GetBridgeDBStorageUnitName() string {

	return this.bridge.GetDBStorageUnitName()
}

func (this *BinaryBridgeRelation[Bridge_Initiator_T, Foreign_Initiator_T]) GetForeignDBStorageUnitName() string {

	return this.foreign.GetDBStorageUnitName()
}

func (this *BinaryBridgeRelation[Bridge_Initiator_T, Foreign_Initiator_T]) ResolveRelation(
	refQueryBuilder relation.IRelationLocalNavigator, foreignInializer relation.IRelationForeignNavigator,
) {

	foreignInializer.Manipulate(
		func(foreign relation.IReadRelationQueryBuilder) {

			foreign.PushRelations(
				this.foreign,
			)
		},
	)
}
