package paginateOutput

import paginateUseCase "app/unitOfWork/genericUsecase/paginate"

type (
	PaginateData[Data_T paginateUseCase.ICursorEntity[interface{}]] struct {
		OptionableQueryPaginateData[Data_T, ignore_optional_query_t]
	}
)
