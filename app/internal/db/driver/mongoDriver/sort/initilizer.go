package sort

import (
	"app/internal/db/driver/mongoDriver/lib"
	"app/internal/db/query"
)

type (
	sort_initializer struct {
		lib.Expression_Map[int]
	}
)

func NewSortInitializer() *sort_initializer {

	return new(sort_initializer)
}

func (this *sort_initializer) Field(name string) query.ISortOperator {

	this.Set(name)

	return this
}

func (this *sort_initializer) Ascending() {

	this.Value(1)
}

func (this *sort_initializer) Descending() {

	this.Value(-1)
}
