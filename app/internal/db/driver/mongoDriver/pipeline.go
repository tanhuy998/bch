package mongoDriver

type (
	mongo_pipeline struct {
		p []interface{}
	}
)

func (this *mongo_pipeline) Add(stages ...interface{}) {

	this.p = append(this.p, stages...)
}
