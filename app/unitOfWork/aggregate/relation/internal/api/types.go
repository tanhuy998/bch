package api

import (
	"app/internal/db/relation"
	"app/internal/db/storage"
)

type (
	IInitializer interface {
		relation.IDBRelationshipDeclarativeInitializer
		storage.IDBStorageIdentifier
	}
)

type (
	IResolver interface {
		relation.IDBRelationResolver
	}
)
