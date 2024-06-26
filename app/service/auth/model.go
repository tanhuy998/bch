package authService

import (
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthClaim string

type AuthUser struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	UserName    string             `json:"username" bson:"uname"`
	Password    []byte             `json:"password" bson:"pw"`
	_isVerified bool
}

type AuthGroup struct {
	UUID    uuid.UUID   `json:"uuid" bson:"uuid"`
	Name    string      `json:"name" bson:"name"`
	UserIds []uuid.UUID `json:"userUUIDs" bson:"userUUIDs"`
}

type AuthField struct {
	UUID     uuid.UUID     `json:"uuid" bson:"uuid"`
	Name     string        `json:"name" bson:"name"`
	Licenses []AuthLicense `json:"licenses" bson:"licenses"`
}

type AuthLicense struct {
	GroupID  uuid.UUID                 `json:"groupID" bson:"groupID"`
	Claims   map[AuthClaim]interface{} `json:"claims" bson:"claims"`
	IssuedAt time.Time                 `json:"issuedAt bson:'issuedAt"`
	Expire   time.Time                 `json:"expire" bson:"expire"`
}

// func (this *User) VerifyPassword(pass string) error {

// 	if this._isVerified {

// 		return fmt.Errorf("")
// 	}

// 	inputBytes := []byte(pass)

// 	err := bcrypt.CompareHashAndPassword(this.Password, inputBytes)

// 	if err != nil {

// 		return err
// 	}

// 	this.Password = nil
// 	return nil
// }

func (this *AuthUser) IsVerified() bool {

	return this._isVerified
}
