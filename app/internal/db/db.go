package db

import (
	"context"
	"errors"
	"fmt"
	"os"
	"regexp"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	CONN_TIMEOUT           = 10
	PING_TIMEOUT           = 3
	QUERY_TIMEOUT          = 5
	ENV_MONGOD_DB_NAME     = "MONGOD_DB_NAME"
	ENV_MONGOD_CONN_STR    = "MONGOD_CONN_STR"
	ENV_MONGOD_CREDENTIAL  = "MONGOD_CREDENTIAL"
	CREDENTIAL_INDENTIFIER = "[username:password@]"
)

var (
	dbClient                  *mongo.Client
	db                        *mongo.Database
	MONGOD_CONN_STR_REGEX     = regexp.MustCompile(`^mongodb(\+srv)?:\/\/\[username:password\]@.*`)
	ERR_DB_CONNECTION_TIMEOUT = errors.New("database connection timeout")
)

type (
	MongoDomainDatabase = mongo.Database
	/*
		for dependency inversion among domain
	*/
	MongoCampaignCollection  = mongo.Collection
	MongoCandidateCollection = mongo.Collection
)

func Init() {

	var err error

	dbClient, err = newClient()

	if err != nil {

		panic(err)
	}

	err = CheckDBConnection()

	if err != nil {

		panic(err)
	}

	dbName := os.Getenv(ENV_MONGOD_DB_NAME)
	db = dbClient.Database(dbName)

	fmt.Println("Initialize database connection")
}

func GetClient() *mongo.Client {

	// if dbClient != nil {

	// 	return dbClient, nil
	// }

	// return newClient()

	return dbClient
}

func CheckDBConnection() error {

	client := GetClient()

	for attempt := 3; attempt > 0; attempt-- {

		ctx, cancel := context.WithTimeout(context.Background(), PING_TIMEOUT+time.Second)
		defer cancel()

		err := client.Ping(ctx, nil)

		if err == nil {

			return nil
		}
	}

	return ERR_DB_CONNECTION_TIMEOUT
}

// func initDB() *mongo.Database {

// 	if db != nil {

// 		return db
// 	}

// 	client, err := GetClient()

// 	if err != nil {

// 		panic(err)
// 	}

// 	dbName := os.Getenv(ENV_MONGOD_DB_NAME)
// 	return client.Database(dbName)
// }

func GetDB() *mongo.Database {

	return db
}

func newClient() (*mongo.Client, error) {

	ctx, cancel := context.WithTimeout(context.Background(), CONN_TIMEOUT*time.Second)
	defer cancel()

	connString := os.Getenv(ENV_MONGOD_CONN_STR) //env.Get(ENV_MONGOD_CONN_STR, "")
	return mongo.Connect(ctx, options.Client().ApplyURI(connString))
}
