package authService

import (
	libCommon "app/src/internal/lib/common"
	"context"
	"fmt"
	"time"

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

// var (
// 	// db.users
// 	lookup_user_claims_on_field bson.D = bson.D{
// 		{"$unwind": "userUUIDs"},
// 		{
// 			"$lookup", bson.M{
// 				"from": ""
// 			}
// 		}
// 	}
// )

type auth_vault struct {
	pending_error  error
	conn_string    string
	auth_db        *mongo.Database
	private_key    []byte
	collectionPool map[string]*mongo.Collection
}

func (this *auth_vault) init() {

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

func (this *auth_vault) initDB() {

	this.initCollections()
	this.initDBIndexes()
}

func (this *auth_vault) initCollections() error {

	if this.pending_error != nil {

		return fmt.Errorf("%s, could not init authentication db collections.", this.pending_error)
	}

	db := this.auth_db
	collNames, err := db.ListCollectionNames(context.TODO(), nil)

	if err != nil {

		this.pending_error = err
	}

	var checkList map[string]*mongo.Collection = map[string]*mongo.Collection{
		USER_COLLECTION:    nil,
		FIELD_COLLECTION:   nil,
		GROUP_COLLECTION:   nil,
		LICENSE_COLLECTION: nil,
	}

	/*
		Check for existence of collections an assign to checklist
	*/
	for _, name := range collNames {

		if _, ok := checkList[name]; !ok {

			continue
		}

		//delete(checkList, name)

		checkList[name] = db.Collection(name)
	}

	/*
		Check for absent collection of the checklist therefore create new one
	*/
	for key, val := range checkList {

		if val != nil {

			continue
		}

		err = db.CreateCollection(context.TODO(), key)

		if err != nil {

			break
		}

		checkList[key] = db.Collection(key)
	}

	if err != nil {

		this.pending_error = err
		return err
	}

	this.collectionPool = checkList

	return nil
}

func (this *auth_vault) getCollection(collName string) (*mongo.Collection, error) {

	ret, ok := this.collectionPool[collName]

	if !ok {

		return nil, fmt.Errorf("Invalid Collection")
	}

	return ret, nil
}

func (this *auth_vault) initDBIndexes() error {

	if this.pending_error != nil {

		return fmt.Errorf("%s, could not initialize db authentication indexes.", this.pending_error)
	}

	db := this.auth_db

	//userColl := db.Collection(USER_COLLECTION)

	_, err := db.Collection(USER_COLLECTION).Indexes().CreateMany(
		context.TODO(),
		[]mongo.IndexModel{
			mongo.IndexModel{
				Keys: "uname",
				Options: &options.IndexOptions{
					Name:   libCommon.PointerPrimitive("username"),
					Unique: libCommon.PointerPrimitive(true),
				},
			},
		},
	)

	if err != nil {

		this.pending_error = err
		return err
	}

	_, err = db.Collection(FIELD_COLLECTION).Indexes().CreateMany(
		context.TODO(),
		[]mongo.IndexModel{
			mongo.IndexModel{
				Keys: "name",
				Options: &options.IndexOptions{
					Name:   libCommon.PointerPrimitive("name"),
					Unique: libCommon.PointerPrimitive(true),
				},
			},
			mongo.IndexModel{
				Keys: "uuid",
				Options: &options.IndexOptions{
					Name:   libCommon.PointerPrimitive("uuid"),
					Unique: libCommon.PointerPrimitive(true),
				},
			},
		},
	)

	if err != nil {

		this.pending_error = err
		return err
	}

	_, err = db.Collection(GROUP_COLLECTION).Indexes().CreateMany(
		context.TODO(),
		[]mongo.IndexModel{
			mongo.IndexModel{
				Keys: "name",
				Options: &options.IndexOptions{
					Name:   libCommon.PointerPrimitive("name"),
					Unique: libCommon.PointerPrimitive(true),
				},
			},
			mongo.IndexModel{
				Keys: "uuid",
				Options: &options.IndexOptions{
					Name:   libCommon.PointerPrimitive("uuid"),
					Unique: libCommon.PointerPrimitive(true),
				},
			},
		},
	)

	if err != nil {

		this.pending_error = err
		return err
	}

	_, err = db.Collection(LICENSE_COLLECTION).Indexes().CreateMany(
		context.TODO(),
		[]mongo.IndexModel{
			mongo.IndexModel{
				Keys: "name",
				Options: &options.IndexOptions{
					Name:   libCommon.PointerPrimitive("name"),
					Unique: libCommon.PointerPrimitive(true),
				},
			},
			mongo.IndexModel{
				Keys: "uuid",
				Options: &options.IndexOptions{
					Name:   libCommon.PointerPrimitive("uuid"),
					Unique: libCommon.PointerPrimitive(true),
				},
			},
		},
	)

	if err != nil {

		this.pending_error = err
		return err
	}

	return nil
}

func (this *auth_vault) getUserByUsername(uname string) (*AuthUser, error) {

	if this.pending_error != nil {

		return nil, this.pending_error
	}

	var user *AuthUser = &AuthUser{
		UserName: uname,
	}

	err := this.retrieveUser(user)

	if err != nil {

		return nil, err
	}

	return user, nil
}

func (this *auth_vault) retrieveUser(user *AuthUser) error {

	if this.pending_error != nil {

		return this.pending_error
	}

	coll, _ := this.getCollection(USER_COLLECTION)

	res := coll.FindOne(context.TODO(), user)
	err := res.Decode(&user)

	if err != nil {

		return err
	}

	return nil
}

// func (this *auth_vault) getUserClaimsOnField(uuid uuid.UUID, field AuthorizationField) {

// 	coll, _ := this.collectionPool[FIELD_COLLECTION]

// 	coll.Aggregate(context.TODO(), lookup_user_claims_on_field)
// }
