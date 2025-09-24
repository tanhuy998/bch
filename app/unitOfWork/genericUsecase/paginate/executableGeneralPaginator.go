package paginateUseCase

import (
	paginateServicePort "app/port/paginate"
	opLog "app/unitOfWork/operationLog"
)

type (
	general_paginator_t[Entity_T any] struct {
		paginator_t[Entity_T, interface{}]
		customPaginator paginateServicePort.IGeneralPaginator
	}
)

func (this general_paginator_t[Entity_T]) GetCursor() interface{} {

	return this.customPaginator.GetCursor()
}

func (this general_paginator_t[Entity_T]) GetPageSize() uint64 {

	return this.paginator_t.Size
}

func (this general_paginator_t[Entity_T]) GetOffset() int64 {

	return this.paginator_t.Offset
}

func (this *general_paginator_t[Entity_T]) Free() {

	this.paginator_t.Free()
}

// func (this *general_paginator_t[Entity_T]) SelfStatistic() repositoryAPI.IStatisticAffectedCountableUnit {

// 	return this.paginator_t.SelfStatistic()
// }

func (this general_paginator_t[Entity_T]) GetLogRecorder() *opLog.LogRecorder {

	return this.paginator_t.GetLogRecorder()
}

func (this general_paginator_t[Entity_T]) get_paginator_initializer() IAbstractOptionableOffsetPaginatorInitializer {

	return &this.paginator_t
}
