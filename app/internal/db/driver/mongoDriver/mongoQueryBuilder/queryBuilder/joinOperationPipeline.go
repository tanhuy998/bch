package queryBuilder

type (
	JoinPipeline struct {
		/*
			this is a projection slice that point to
			this.MongoAggregateQueryBuilder.mongo_pipeline.P in order to parse
			the pointed slice as $lookup pipeline.
		*/
		Pipeline                   []interface{} `json:"pipeline,omitempty" bson:"pipeline,omitempty"`
		MongoAggregateQueryBuilder `json:",inline" bson:"-"`
	}
)

func (this *JoinPipeline) Init() {

	this.initJoinPipeline()
}

func (this *JoinPipeline) initJoinPipeline() {

}

func (this *JoinPipeline) Done() {

	this.MongoAggregateQueryBuilder.Done()

	if len(this.MongoAggregateQueryBuilder.P) == 0 {

		this.Pipeline = nil
		return
	}

	this.Pipeline = this.MongoAggregateQueryBuilder.P
}
