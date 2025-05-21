package paginateOutput

import paginateUseCase "app/unitOfWork/genericUsecase/paginate"

type (
	OptionalQueryPaginateData[
		Data_T paginateUseCase.ICursorEntity[interface{}],
		Option_Query_T comparable,
	] struct {
		OptionableQueryPaginateData[Data_T, Option_Query_T]
	}
)

func (this *OptionalQueryPaginateData[Data_T, Option_Query_T]) IgnoreOptionQuery() {

	this.ignoreOptionalQuery = true

	this.unlinkOptionQuery()
}
