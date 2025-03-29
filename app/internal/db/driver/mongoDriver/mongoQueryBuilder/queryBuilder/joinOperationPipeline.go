package queryBuilder

type (
	JoinPipeline struct {
		/*
			this is a projection slice that point to
			this.MongoAggregateQueryBuilder.mongo_pipeline.P in order to parse
			the pointed slice as $lookup pipeline.
		*/
		Pipeline *[]interface{} `json:"pipeline,omitempty" bson:"pipeline,omitempty"`
		// MongoAggregateQueryBuilder `json:",inline" bson:"-"` // $lookup stage's pipeline, struct metadatas for bson are defined in mongo_pipeline struct
		MongoAggregateQueryBuilder `json:",inline" bson:"-"`
	}
)

func (this *JoinPipeline) Init() {

	this.initJoinPipeline()
}

func (this *JoinPipeline) initJoinPipeline() {

	if this.Pipeline != nil {

		return
	}

	// if len(this.MongoAggregateQueryBuilder.P) == 0 {

	// 	this.MongoAggregateQueryBuilder.mongo_pipeline.init()
	// }

	// this.mongo_pipeline.init()

	// this.Pipeline = &this.MongoAggregateQueryBuilder.mongo_pipeline.P

	//this.Pipeline = this.MongoAggregateQueryBuilder.GetRefPipeline()
}

// func (this *join_pipeline) SetPipeline(ref IJoinPipeline) {

// 	this.IJoinPipeline = ref
// }

// func (this *join_pipeline) getPipelineQueryBuilder() *MongoAggregateQueryBuilder {

// 	return &this.MongoAggregateQueryBuilder
// }

func (this *JoinPipeline) Done() {

	// this.MongoAggregateQueryBuilder.Done()

	this.MongoAggregateQueryBuilder.Done()

	switch {
	case this.Pipeline == nil:
		return
	case *this.Pipeline == nil:
		return
	case len(*this.Pipeline) == 0:
		this.Pipeline = nil
	}
}
