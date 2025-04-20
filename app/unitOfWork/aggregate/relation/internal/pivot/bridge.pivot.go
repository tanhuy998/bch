package pivot

import "app/internal/db/storage"

type (
	Bridge[Repo_Entity_T any] struct {
		Pivot storage.IDBStoragePivot[Repo_Entity_T]
	}
)

func (this Bridge[Repo_Entity_T]) GetBridgeDBStorageUnitName() string {

	return this.Pivot.GetDBStorageUnitName()
}
