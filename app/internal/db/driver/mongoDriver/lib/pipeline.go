package lib

import (
	"app/internal/db/storage"
)

type (
	IGenericPipeline interface {
		storage.IArbitraryQuery
		PushStages(stages ...interface{})
		PrependStages(stages ...interface{})
	}

	IClonableGenericPipeline[Concrete_Pipeline any] interface {
		IGenericPipeline
		Clone() *Concrete_Pipeline
	}
)

type (
	MongoPipeline struct {
		P []interface{} `json:"debug_pipeline" bson:"-"`
	}
)

func (this *MongoPipeline) init() {

	if this.P != nil {

		return
	}

	this.P = make([]interface{}, 0)
}

func (this *MongoPipeline) PushStages(stages ...interface{}) {

	if this.P == nil {

		this.P = stages
		return
	}

	this.P = append(this.P, stages...)
}

func (this *MongoPipeline) PrependStages(stages ...interface{}) {

	if this.P == nil {

		this.P = stages
		return
	}

	this.P = append(stages, this.P...)
}

func (this *MongoPipeline) GetArbitraryQuery() interface{} {

	return this.P[:]
}
