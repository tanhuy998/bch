package bridgeRelation

import "app/internal/db/query"

type (
	IForeignDeclarativeInitializer interface {
		GetForeignInitFunc() query.JoinInitFunc
	}

	IBridgeDeclarativeInitializer interface {
		GetBridgeInitFunc() query.JoinInitFunc
	}
)

type (
	IBridgeStorageUnitIdentifier interface {
		GetBridgeDBStorageUnitName() string
	}

	IForeignStorageUnitIdentifier interface {
		GetForeignDBStorageUnitName() string
	}
)

type (
	IGeneralBinaryRelationDeclarativeInitializer interface {
		IForeignDeclarativeInitializer
		IBridgeDeclarativeInitializer
		IBridgeStorageUnitIdentifier
		IForeignStorageUnitIdentifier
	}
)
