package storage

import (
	"context"
)

type (
	/*
		IDBStorageUnit interface defines a way to get the object which is mapped
		to a specific storage unit (database's table/collection).
	*/
	IDBStorageUnit[DBStorage_T, Model_T any] interface {
		IDBStorageIdentifier
		IDBStorageUnitGetter[DBStorage_T]
	}

	IDBStorageUnitGetter[DBStorage_T any] interface {
		GetStorageUnit() *DBStorage_T
	}

	IDBStorageIdentifier interface {
		GetDBStorageUnitName() string
	}

	// This interface is used for detemining which repository whose storage unit refer to the target
	// collection/table.
	IDBStorageQueryExecutor[Local_Storage_Unit_Entity_T any] interface {
		ToSlice(resulSlice interface{}, query IArbitraryQuery, ctx context.Context) error
		First(result interface{}, query IArbitraryQuery, ctx context.Context) error
	}
)
