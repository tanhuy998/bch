package requestInput

import (
	"app/valueObject"
)

type (
	AuthorityInput struct {
		valueObject.IAuthorityData
	}
)

func (this *AuthorityInput) GetAuthority() valueObject.IAuthorityData {

	return this.IAuthorityData
}

func (this *AuthorityInput) SetAuthority(auth valueObject.IAuthorityData) {

	this.IAuthorityData = auth
}
