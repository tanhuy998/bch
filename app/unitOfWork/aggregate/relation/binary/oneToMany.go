package binaryRelation

// import (
// 	"app/unitOfWork/aggregate/relation/binary/api"
// 	"app/unitOfWork/aggregate/relation/internal"
// 	"app/unitOfWork/aggregate/relation/internal/initializer"
// 	"app/unitOfWork/aggregate/relation/internal/resolver"
// )

// type (
// 	OneToManyBy[
// 		Relation_T api.IGeneralBinaryRelationDeclarativeInitializer,
// 	] struct {
// 		BinaryBridgeRelationInitiator[
// 			internal.Initiator[
// 				resolver.One_To_One,
// 				initializer.Bridge[Relation_T],
// 			],
// 			internal.Initiator[
// 				resolver.One_To_Many,
// 				initializer.Foreign[Relation_T],
// 			],
// 		]
// 	}
// )

// func (this OneToManyBy[Relation_T]) GetDBRelationKind() string {

// 	return "bridge_one_to_many"
// }
