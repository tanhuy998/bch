package repository

import (
	dbQueryTracerPort "app/port/dbQueryTracer"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	IQueryTracer interface {
		Trace(collectionName string, label string, ctx context.Context) (stop func(error))
	}

	MongoDBQueryMonitorCollection struct {
		collection *mongo.Collection
		Tracer     dbQueryTracerPort.IDBQueryTracer
	}
)

func (this *MongoDBQueryMonitorCollection) SetTracer(t dbQueryTracerPort.IDBQueryTracer) {

	if t == nil {

		panic("MongoDBQueryMonitorCollection error: tracer must not be nil")
	}

	this.Tracer = t
}

func (this *MongoDBQueryMonitorCollection) BulkWrite(
	ctx context.Context,
	models []mongo.WriteModel,
	opts ...*options.BulkWriteOptions,
) (ret *mongo.BulkWriteResult, err error) {

	defer this.Tracer.Trace(this.collection.Name(), "buld_write", ctx)(err)

	ret, err = this.collection.BulkWrite(ctx, models, opts...)
	return
}

func (this *MongoDBQueryMonitorCollection) InsertOne(
	ctx context.Context,
	document interface{},
	opts ...*options.InsertOneOptions,
) (ret *mongo.InsertOneResult, err error) {

	defer this.Tracer.Trace(this.collection.Name(), "insert_one", ctx)(err)

	ret, err = this.collection.InsertOne(ctx, document, opts...)

	return
}

func (this *MongoDBQueryMonitorCollection) InsertMany(
	ctx context.Context,
	documents []interface{},
	opts ...*options.InsertManyOptions,
) (ret *mongo.InsertManyResult, err error) {

	defer this.Tracer.Trace(this.collection.Name(), "insert_many", ctx)(err)

	ret, err = this.collection.InsertMany(ctx, documents, opts...)

	return
}

func (this *MongoDBQueryMonitorCollection) DeleteOne(
	ctx context.Context,
	filter interface{},
	opts ...*options.DeleteOptions,
) (ret *mongo.DeleteResult, err error) {

	defer this.Tracer.Trace(this.collection.Name(), "delete_one", ctx)(err)

	ret, err = this.collection.DeleteOne(ctx, filter, opts...)

	return
}

func (this *MongoDBQueryMonitorCollection) DeleteMany(
	ctx context.Context,
	filter interface{},
	opts ...*options.DeleteOptions,
) (ret *mongo.DeleteResult, err error) {

	defer this.Tracer.Trace(this.collection.Name(), "delete_many", ctx)(err)

	ret, err = this.collection.DeleteMany(ctx, filter, opts...)

	return
}

func (this *MongoDBQueryMonitorCollection) UpdateByID(
	ctx context.Context,
	id interface{},
	update interface{},
	opts ...*options.UpdateOptions,
) (ret *mongo.UpdateResult, err error) {

	defer this.Tracer.Trace(this.collection.Name(), "update_by_id", ctx)(err)

	ret, err = this.collection.UpdateByID(ctx, id, update, opts...)

	return
}

func (this *MongoDBQueryMonitorCollection) UpdateOne(
	ctx context.Context,
	filter interface{},
	update interface{},
	opts ...*options.UpdateOptions,
) (ret *mongo.UpdateResult, err error) {

	defer this.Tracer.Trace(this.collection.Name(), "update_one", ctx)(err)

	ret, err = this.collection.UpdateOne(ctx, filter, update, opts...)

	return
}

func (this *MongoDBQueryMonitorCollection) UpdateMany(
	ctx context.Context,
	filter interface{},
	update interface{},
	opts ...*options.UpdateOptions,
) (ret *mongo.UpdateResult, err error) {

	defer this.Tracer.Trace(this.collection.Name(), "update_many", ctx)(err)

	ret, err = this.collection.UpdateMany(ctx, filter, update, opts...)

	return
}

func (this *MongoDBQueryMonitorCollection) ReplaceOne(
	ctx context.Context,
	filter interface{},
	replacement interface{},
	opts ...*options.ReplaceOptions,
) (ret *mongo.UpdateResult, err error) {

	defer this.Tracer.Trace(this.collection.Name(), "replace_one", ctx)(err)

	ret, err = this.collection.ReplaceOne(ctx, filter, replacement, opts...)

	return
}

func (this *MongoDBQueryMonitorCollection) Aggregate(
	ctx context.Context,
	pipeline interface{},
	opts ...*options.AggregateOptions,
) (ret *mongo.Cursor, err error) {

	defer this.Tracer.Trace(this.collection.Name(), "aggregate", ctx)(err)

	ret, err = this.collection.Aggregate(ctx, pipeline, opts...)

	return
}

func (this *MongoDBQueryMonitorCollection) Distinct(
	ctx context.Context,
	fieldName string,
	filter interface{},
	opts ...*options.DistinctOptions,
) (ret []interface{}, err error) {

	defer this.Tracer.Trace(this.collection.Name(), "distinct", ctx)(err)

	ret, err = this.collection.Distinct(ctx, fieldName, filter, opts...)

	return
}

func (this *MongoDBQueryMonitorCollection) Find(
	ctx context.Context,
	filter interface{},
	opts ...*options.FindOptions,
) (ret *mongo.Cursor, err error) {

	defer this.Tracer.Trace(this.collection.Name(), "find", ctx)(err)

	ret, err = this.collection.Find(ctx, filter, opts...)

	return
}

func (this *MongoDBQueryMonitorCollection) FindOne(
	ctx context.Context,
	filter interface{},
	opts ...*options.FindOneOptions,
) *mongo.SingleResult {

	stopTrace := this.Tracer.Trace(this.collection.Name(), "find_one", ctx)

	res := this.collection.FindOne(ctx, filter, opts...)

	stopTrace(res.Err())

	return res
}

func (this *MongoDBQueryMonitorCollection) FindOneAndDelete(
	ctx context.Context,
	filter interface{},
	opts ...*options.FindOneAndDeleteOptions,
) *mongo.SingleResult {

	stopTrace := this.Tracer.Trace(this.collection.Name(), "find_one_and_delete", ctx)

	res := this.collection.FindOneAndDelete(ctx, filter, opts...)

	stopTrace(res.Err())

	return res
}

func (this *MongoDBQueryMonitorCollection) FindOneAndReplace(
	ctx context.Context,
	filter interface{},
	replacement interface{},
	opts ...*options.FindOneAndReplaceOptions,
) *mongo.SingleResult {

	stopTrace := this.Tracer.Trace(this.collection.Name(), "find_one_and_replace", ctx)

	res := this.collection.FindOneAndReplace(ctx, filter, replacement, opts...)

	stopTrace(res.Err())

	return res
}

func (this *MongoDBQueryMonitorCollection) FindOneAndUpdate(
	ctx context.Context,
	filter interface{},
	update interface{},
	opts ...*options.FindOneAndUpdateOptions,
) *mongo.SingleResult {

	stopTrace := this.Tracer.Trace(this.collection.Name(), "find_one_and_update", ctx)

	res := this.collection.FindOneAndUpdate(ctx, filter, update, opts...)

	stopTrace(res.Err())

	return res
}

func (this *MongoDBQueryMonitorCollection) Database() *mongo.Database {

	return this.collection.Database()
}

func (this *MongoDBQueryMonitorCollection) CountDocuments(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error) {

	return this.collection.CountDocuments(
		ctx, filter, opts...,
	)
}

func (this *MongoDBQueryMonitorCollection) EstimatedDocumentCount(ctx context.Context, opts ...*options.EstimatedDocumentCountOptions) (int64, error) {

	return this.collection.EstimatedDocumentCount(
		ctx, opts...,
	)
}

func (this *MongoDBQueryMonitorCollection) Watch(ctx context.Context, pipeline interface{}, opts ...*options.ChangeStreamOptions) (*mongo.ChangeStream, error) {

	return this.collection.Watch(
		ctx, pipeline, opts...,
	)
}
