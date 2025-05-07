package requestInput

import (
	"app/valueObject"
)

type (
	AuthorityInput struct {
		Auth valueObject.IAuthorityData
	}
)

func (this *AuthorityInput) GetAuthority() valueObject.IAuthorityData {

	return this.Auth //.IAuthorityData
}

func (this *AuthorityInput) SetAuthority(auth valueObject.IAuthorityData) {

	// this.IAuthorityData = auth
	this.Auth = auth
}

func (this *AuthorityInput) HasAuthority() bool {

	// return this.IAuthorityData != nil
	return this.Auth != nil
}
