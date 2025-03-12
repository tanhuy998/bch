package mongoDriver

// import (
// 	"app/internal/db/driver/mongoDriver/mongoQueryBuilder"
// 	"app/internal/db/driver/mongoDriver/mongoStorage"
// 	accessLogServicePort "app/port/accessLog"
// 	"context"

// 	"go.mongodb.org/mongo-driver/bson"
// )

// type (
// 	QueryExecutor[Entity_T any] struct {
// 		AccessLogger accessLogServicePort.IAccessLogger
// 		delegator    mongoStorage.QueryExecutorProxy[Entity_T] //lib.IMongoDBCollection //*MongoDBDelegator[Model_T]
// 		mongoQueryBuilder.MongoAggregateQueryBuilder
// 	}
// )

// func (this *QueryExecutor[Entity_T]) exec(ctx context.Context) ([]Entity_T, error) {

// 	// return lib.Aggregate[Entity_T](
// 	// 	this.delegator,
// 	// 	this.GetPipeline(), //this.mongo_pipeline.p,
// 	// 	ctx,
// 	// )

// 	return this.delegator.Exec(
// 		this.GetPipeline(),
// 		ctx,
// 	)
// }

// func (this *QueryExecutor[Model_T]) First(ctx context.Context) (*Model_T, error) {

// 	this.PushStages(
// 		bson.D{
// 			{"$limit", 1},
// 		},
// 	)

// 	// return lib.AggregateOne[Model_T](
// 	// 	this.delegator,
// 	// 	this.GetPipeline(), //this.mongo_pipeline.p,
// 	// 	ctx,
// 	// )

// 	return this.delegator.First(
// 		this.GetPipeline(),
// 		ctx,
// 	)
// }

// func (this *QueryExecutor[Model_T]) All(ctx context.Context) ([]Model_T, error) {

// 	return this.exec(ctx)
// }
