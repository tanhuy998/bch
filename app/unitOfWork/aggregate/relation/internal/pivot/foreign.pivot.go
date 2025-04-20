package pivot

import "app/internal/db/storage"

type (
	Foreign[Repo_Entity_T any] struct {
		Pivot storage.IDBStoragePivot[Repo_Entity_T]
	}
)

func (this Foreign[Repo_Entity_T]) GetForeignDBStorageUnitName() string {

	return this.Pivot.GetDBStorageUnitName()
}
