package binding

import (
	"app/internal/db/driver/mongoDriver/mongoStorage"
	"app/internal/db/storage"
	repositoryAPI "app/repository/api"
	irisIoc "app/shared/lib/iris/ioc"
	iocOption "app/shared/lib/iris/ioc/option"

	"github.com/kataras/iris/v12/hero"
)

type (
	RepoRegisterFunc func(container *hero.Container)
)

func ByRepositoryOf[Entity_T any](
	concreteRepoObj repositoryAPI.ICRUDRepository[Entity_T],
) RepoRegisterFunc {

	return func(container *hero.Container) {

		irisIoc.RegisterArbitrary(
			container, concreteRepoObj,
			iocOption.StructDependents(true),
			iocOption.BindAs[storage.IDBStoragePivot[Entity_T]](),
			iocOption.BindAs[mongoStorage.IMongoDBStorageUnit[Entity_T]](),
			iocOption.BindAs[repositoryAPI.ICRUDRepository[Entity_T]](),
			iocOption.BindAs[storage.IDBStorageQueryExecutor[Entity_T]](),
		)
	}
}
