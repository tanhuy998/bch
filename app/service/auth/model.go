package authService

import "github.com/google/uuid"

type AuthClaim string

type AuthUser struct {
	UserName    string `json:"username" bson:"uname"`
	Password    []byte `json:"password" bson:"pw"`
	_isVerified bool
}

type AuthGroup struct {
	UUID   uuid.UUID   `json:"uuid" bson:"uuid"`
	Name   string      `json:"name" bson:"name"`
	UserId []uuid.UUID `json:"userIDs" bson:"userIDs"`
}

type AuthField struct {
	UUID          uuid.UUID   `json:"uuid" bson:"uuid"`
	Name          string      `json:"name" bson:"name"`
	GrantedGroups []AuthGroup `json:"grantedGroups" bson:"grantedGroups"`
}

type FieldPermission struct {
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
