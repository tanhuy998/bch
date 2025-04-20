package binaryRelation

// type (
// 	IBridge interface {
// 		relation.IDBRelationshipDeclarativeInitializer
// 	}
// )

// type (
// 	OneToOneBy[
// 		Relation_T api.IGeneralBinaryRelationDeclarativeInitializer,
// 	] struct {
// 		BinaryBridgeRelationInitiator[
// 			internal.Initiator[
// 				resolver.One_To_One,
// 				initializer.Bridge[Relation_T],
// 			],
// 			internal.Initiator[
// 				resolver.One_To_One,
// 				initializer.Foreign[Relation_T],
// 			],
// 		]
// 	}
// )

// type (
// 	OneToOneBy[
// 		Bridge_Entity_T any,
// 		Foreign_Entity_T any,
// 	] struct {
// 		BinaryBridgeRelationInitiator[
// 			bridge.OneToOne[Bridge_Entity_T],
// 			foreign.OneToOne[Foreign_Entity_T],
// 		]
// 	}
// )

// func (this OneToOneBy[Relation_T]) GetDBRelationKind() string {

// 	return "bridge_one_to_one"
// }

// func (this *OneToOneBy[Relation_T]) GetForeignAlias() string {

// 	return this.rel.GetBridgeDBStorageUnitName()
// }

// func (this *OneToOneBy[Relation_T]) ResolveRelation(
// 	local relation.IRelationLocalNavigator, foreignInializer relation.IRelationForeignNavigator,
// ) {

// 	this.AbstractBridgeRelation._resolveRelation(
// 		local, foreignInializer,
// 	)

// 	local.Join()

// 	foreignInializer.Manipulate(
// 		func(foreign relation.IReadRelationQueryBuilder) {

// 			foreign.PushRelations(
// 				this.foreign_relation_initiator,
// 			)
// 		},
// 	)

// }

// func (this *OneToOneBy[Relation_T]) GetDBRelationKind() string {

// 	return "bridge_one_to_one"
// }

// func (this *OneToOneBy[Relation_T]) GetRelationInitFunc() query.JoinInitFunc {

// 	return this.rel.GetBridgeInitFunc()
// }
