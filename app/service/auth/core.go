package authService

import (
	"context"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	}

	this.auth_db = client.Database("auth")
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
