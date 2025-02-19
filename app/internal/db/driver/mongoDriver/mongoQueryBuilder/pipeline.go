package mongoQueryBuilder

type (
	mongo_pipeline struct {
		p []interface{}
	}
)

func (this *mongo_pipeline) PushStages(stages ...interface{}) {

	this.p = append(this.p, stages...)
}
