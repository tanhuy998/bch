package authService

import (
	"context"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DB_NAME            = "auth"
	USER_COLLECTION    = "users"
	FIELD_COLLECTION   = "fields"
	GROUP_COLLECTION   = "groups"
	LICENSE_COLLECTION = "licenses"
	COLL_NUM           = 4
)

/*
Authentication service is design in microservice manner in order to isolate the authentication
and authorization logics with the app's bussiness logic
in inplementing Policy Decision Point pattern.

Authorization implement the following principal
claim-based authorization
group-based authorization
document-based authorization
*/

type auth_core struct {
	IAuthCore
	pending_error error
	conn_string   string
	auth_db       *mongo.Database
	private_key   []byte
}

func (this *auth_core) init() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, error := mongo.Connect(ctx, options.Client().ApplyURI(this.conn_string))

	if error != nil {

		this.pending_error = error

		return
	}

	this.auth_db = client.Database(DB_NAME)

	this.initDB()
}

func (this *auth_core) ValidateCredential(uname string, pass string) (string, error) {

	if this.pending_error != nil {

		return "", this.pending_error
	}

	return this.generateToken([16]byte{})
}

func (this *auth_core) AuthorizeClaims(token string, field string, claims []string) error {

	if this.pending_error != nil {

		return this.pending_error
	}

	return nil
}

func (this *auth_core) AuthorizeGroup(token string, field string, groups []string) error {

	if this.pending_error != nil {

		return this.pending_error
	}

	return nil
}

func (this *auth_core) generateToken(uuid uuid.UUID) string {

	return ""
}

func (this *auth_core) initDB() {

	this.checkCollection()
}

func (this *auth_core) checkCollection() error {

	if this.pending_error != nil {

		return this.pending_error
	}

	db := this.auth_db
	collNames, err := db.ListCollectionNames(context.TODO(), nil)

	if err != nil {

		this.pending_error = err
	}

	var checkList map[string]bool = map[string]bool{
		USER_COLLECTION:    false,
		FIELD_COLLECTION:   false,
		GROUP_COLLECTION:   false,
		LICENSE_COLLECTION: false,
	}

	for _, name := range collNames {

		if _, ok := checkList[name]; !ok {

			continue
		}

		delete(checkList, name)
	}

	for key := range checkList {

		db.CreateCollection(context.TODO(), key)
	}

	return nil
}
