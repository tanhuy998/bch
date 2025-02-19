package storage

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
)
