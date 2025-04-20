package api

import (
	"app/internal/db/query"
	"app/internal/db/storage"
)

type (
	IForeignDeclarativeInitializer interface {
		GetForeignInitFunc() query.JoinInitFunc
	}

	IBridgeDeclarativeInitializer interface {
		GetBridgeInitFunc() query.JoinInitFunc
	}

	IBridgeStoragePivot[Entity_T any] interface {
		GetBridgeStoragePivot() storage.IDBStoragePivot[Entity_T]
	}

	IForeign interface {
		IForeignDeclarativeInitializer
		IForeignStorageUnitIdentifier
	}
)

type (
	IBridgeStorageUnitIdentifier interface {
		GetBridgeDBStorageUnitName() string
	}

	IForeignStorageUnitIdentifier interface {
		GetForeignDBStorageUnitName() string
	}

	IForeignStoragePivot[Entity_T any] interface {
		GetForeignStoragePivot() storage.IDBStoragePivot[Entity_T]
	}

	IBridge interface {
		IBridgeDeclarativeInitializer
		IBridgeStorageUnitIdentifier
	}
)

type (
	IGeneralBinaryRelationDeclarativeInitializer interface {
		IBridge
		IForeign
	}
)
