package binaryRelationFacade

import (
	"app/internal/db/query"
	"app/internal/db/relation"

	binaryRelation "app/unitOfWork/aggregate/relation/binary"
	binaryAPI "app/unitOfWork/aggregate/relation/binary/api"
	"app/unitOfWork/aggregate/relation/internal"
	"app/unitOfWork/aggregate/relation/internal/api"
	"app/unitOfWork/aggregate/relation/internal/closure"
	"app/unitOfWork/aggregate/relation/internal/initializer"
)

type (
	IBinaryRelationTriat interface {
		binaryAPI.IGeneralBinaryRelationDeclarativeInitializer
		GetBridgeResolver() relation.IDBRelationResolver
		GetForeignResolver() relation.IDBRelationResolver
	}
)

func Marshall[Initializer_T IBinaryRelationTriat](
	RelationInitiator Initializer_T,
) relation.IDBRelationInitiator {

	return marshall_binary_relation(
		resolve_initiators(RelationInitiator),
	)
}

func resolve_initiators[Initializer_T IBinaryRelationTriat](
	RelationInitiator Initializer_T,
) (bridge relation.IDBRelationInitiator, foreign relation.IDBRelationInitiator) {

	bridge = internal.Initiator[
		api.IResolver, initializer.Bridge[Initializer_T],
	]{
		Resolver: RelationInitiator.GetBridgeResolver(),
		Initiator: initializer.Bridge[Initializer_T]{
			RelationInitiator,
		},
	}

	foreign = internal.Initiator[
		api.IResolver, initializer.Foreign[Initializer_T],
	]{
		Resolver: RelationInitiator.GetForeignResolver(),
		Initiator: initializer.Foreign[Initializer_T]{
			RelationInitiator,
		},
	}

	return
}

func marshall_binary_relation(
	bridge relation.IDBRelationInitiator, foreign relation.IDBRelationInitiator,
) relation.IDBRelationInitiator {

	return binaryRelation.BinaryBridgeRelationInitiator[
		relation.IDBRelationInitiator, relation.IDBRelationInitiator,
	]{
		bridge, foreign,
	}
}

func MarshallWithFilter[Initializer_T IBinaryRelationTriat](
	RelationInitiator Initializer_T, bridgeFilterFunc query.FilterFunc, foreignFilterFunc query.FilterFunc,
) relation.IDBRelationInitiator {

	bridge, foreign := resolve_initiators(RelationInitiator)

	if bridgeFilterFunc != nil {

		bridge = closure.NewFilterInitiator(bridge, bridgeFilterFunc)
	}

	if foreignFilterFunc != nil {

		foreign = closure.NewFilterInitiator(foreign, bridgeFilterFunc)
	}

	return marshall_binary_relation(
		bridge, foreign,
	)
}

func MarshallWithConditions[Initializer_T IBinaryRelationTriat](
	RelationInitiator Initializer_T,
	BridgeConditionFunc query.DataConditionMatchFunc,
	ForeignConditionFunc query.DataConditionMatchFunc,
) relation.IDBRelationInitiator {

	bridge, foreign := resolve_initiators(RelationInitiator)

	return marshall_binary_relation(
		bridge, foreign,
	)
}
