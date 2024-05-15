package db

import (
	"app/app/bootstrap"
	"context"
	"regexp"
	"time"

	"github.com/gofor-little/env"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	CONN_TIMEOUT           = 10
	QUERY_TIMEOUT          = 5
	ENV_MONGOD_DB_NAME     = "MONOGOD_DB_NAME"
	ENV_MONGOD_CONN_STR    = "MONGOD_CONN_STR"
	ENV_MONGOD_CREDENTIAL  = "MONGOD_CREDENTIAL"
	CREDENTIAL_INDENTIFIER = "[username:password@]"
)

var (
	dbClient                 *mongo.Client
	db                       *mongo.Database
	MONGOD_CONN_STR_REGEX, _ = regexp.Compile(`^mongodb(\+srv)?:\/\/\[username:password\]@.*`)
)

func GetClient() (*mongo.Client, error) {

	if dbClient != nil {

		return dbClient, nil
	}

	return newClient()
}

func GetDB() *mongo.Database {

	if db != nil {

		return db
	}

	client, err := GetClient()

	if err != nil {

		panic(err)
	}

	dbName := env.Get(ENV_MONGOD_DB_NAME, "dev")
	return client.Database(dbName)
}

func newClient() (*mongo.Client, error) {

	err := bootstrap.InitEnv()

	if err != nil {

		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), CONN_TIMEOUT*time.Second)
	defer cancel()

	connString := env.Get(ENV_MONGOD_CONN_STR, "")
	return mongo.Connect(ctx, options.Client().ApplyURI(connString))
}
