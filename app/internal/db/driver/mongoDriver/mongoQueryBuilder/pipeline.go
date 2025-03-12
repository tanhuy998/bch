package mongoQueryBuilder

type (
	mongo_pipeline struct {
		P []interface{} //`json:"pipeline,omitempty" bson:"pipeline,omitempty"`
	}
)

func (this *mongo_pipeline) init() {

	if len(this.P) > 0 {

		return
	}

	this.P = make([]interface{}, 0)
}

func (this *mongo_pipeline) PushStages(stages ...interface{}) {

	if this.P == nil {

		this.P = stages
		return
	}

	this.P = append(this.P, stages...)
}

func (this *mongo_pipeline) PrependStages(stages ...interface{}) {

	if this.P == nil {

		this.P = stages
		return
	}

	this.P = append(stages, this.P...)
}

func (this *mongo_pipeline) GetArbitraryQuery() interface{} {

	return this.P[:]
}
