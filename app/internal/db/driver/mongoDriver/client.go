package mongoDriver

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type (
	MongoClient struct {
		Client *mongo.Client
	}
)

func (this *MongoClient) Transaction(initCtx context.Context, fn func(context.Context) (interface{}, error)) (interface{}, error) {

	session, err := this.Client.StartSession()

	if err != nil {

		return nil, err
	}

	defer session.EndSession(initCtx)

	return session.WithTransaction(
		initCtx,
		func(ctx mongo.SessionContext) (interface{}, error) {

			return fn(ctx)
		},
	)
}
