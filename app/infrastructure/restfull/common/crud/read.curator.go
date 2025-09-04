package crud

import (
	"app/infrastructure/restfull/common/endpoint"
)

type (
	/*
		Types that embed this class must not override BeforeActivation(mvc.BeforeActivation) method
	*/
	ReadEndpointCurator struct {
		endpoint.APIEndpointCurator
	}
)

func (this *ReadEndpointCurator) GET(path string) endpoint.IEndpointBuilder {

	return endpoint.Handle(&this.APIEndpointCurator, "GET", path)
}

func (this *ReadEndpointCurator) HEAD(path string) endpoint.IEndpointBuilder {

	return endpoint.Handle(&this.APIEndpointCurator, "HEAD", path)
}
