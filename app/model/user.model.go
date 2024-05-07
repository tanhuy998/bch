package model

import (
	"app/app/lib/abstract"
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	abstract.IAuthenticate
	abstract.IAuthorize
	UUID        uuid.UUID       `json:"uuid" bson:"uuid"`
	UserName    string          `json:"userName" bson:"userName"`
	Password    []byte          `json:"password" bson:"password"`
	Roles       []abstract.Role `json:"roles" bson:"roles"`
	_isVerified bool
}

type UserGroup struct {
	UUID       uuid.UUID   `json:"uuid" bson:"uuid"`
	User_UUIDs []uuid.UUID `json:"users_uuid" bson:"users_uuid"`
}

func (this *User) VerifyPassword(pass string) error {

	if this._isVerified {

		return fmt.Errorf("")
	}

	inputBytes := []byte(pass)

	err := bcrypt.CompareHashAndPassword(this.Password, inputBytes)

	if err != nil {

		return err
	}

	this.Password = nil
	return nil
}

func (this *User) IsVerified() bool {

	return this._isVerified
}

func (this *User) GetRoles() []abstract.Role {

	return []abstract.Role{}
}

func (this *User) Authorize(required []abstract.Role) error {

	return nil
}
