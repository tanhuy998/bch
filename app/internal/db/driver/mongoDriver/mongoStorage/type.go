package mongoStorage

import "app/internal/db/storage"

type (
	IMongoDBStorageUnit[Source_Repo_Entity_T any] interface {
		storage.IDBStorageUnit[MongoDBQueryMonitorCollection, Source_Repo_Entity_T]
	}
)
