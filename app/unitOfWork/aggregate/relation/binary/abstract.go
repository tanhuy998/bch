package binaryRelation

import (
	"app/internal/db/query"
	"app/internal/db/relation"
	"fmt"
)

type (
	BinaryBridgeRelationInitiator[
		Bridge_Initiator_T relation.IDBRelationInitiator,
		Foreign_Initiator_T relation.IDBRelationInitiator,
	] struct {
		Bridge  Bridge_Initiator_T
		Foreign Foreign_Initiator_T
	}
)

/*
Implemnt
*/
func (this BinaryBridgeRelationInitiator[Bridge_Initiator_T, Foreign_Initiator_T]) GetRelationInitFunc() query.JoinInitFunc {

	defer func() {

		if r := recover(); r != nil {

			panic(fmt.Sprintf("BinaryBridgeRelation panic, %s", r))
		}
	}()

	return this.Bridge.GetRelationInitFunc()
}

// func (this BinaryBridgeRelationInitiator[Bridge_Initiator_T, Foreign_Initiator_T]) GetBridgeDBStorageUnitName() string {

// 	return this.bridge.GetDBStorageUnitName()
// }

// func (this BinaryBridgeRelationInitiator[Bridge_Initiator_T, Foreign_Initiator_T]) GetForeignDBStorageUnitName() string {

// 	return this.foreign.GetDBStorageUnitName()
// }

func (this BinaryBridgeRelationInitiator[Bridge_Initiator_T, Foreign_Initiator_T]) GetDBStorageUnitName() string {

	return this.Bridge.GetDBStorageUnitName()
}

func (this BinaryBridgeRelationInitiator[Bridge_Initiator_T, Foreign_Initiator_T]) ResolveRelation(
	refQueryBuilder relation.IRelationLocalNavigator, foreignInializer relation.IRelationForeignNavigator,
) {

	foreignInializer.Manipulate(
		func(foreign relation.IReadRelationQueryBuilder) {

			foreign.PushRelations(
				this.Foreign,
			)
		},
	)
}

func (this BinaryBridgeRelationInitiator[Bridge_Initiator_T, Foreign_Initiator_T]) GetDBRelationKind() string {

	return "binary_relation"
}
