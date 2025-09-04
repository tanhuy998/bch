package crud

import (
	"app/infrastructure/restfull/common/endpoint"
)

type (
	/*
		Types that embed this class must not override BeforeActivation(mvc.BeforeActivation) method
	*/
	ReadEndpointBuilder struct {
		endpoint.EndpointBuilder
	}
)

func (this *ReadEndpointBuilder) GET(path string) endpoint.IEndpointInitiator {

	return endpoint.Handle(&this.EndpointBuilder, "GET", path)
}

func (this *ReadEndpointBuilder) HEAD(path string) endpoint.IEndpointInitiator {

	return endpoint.Handle(&this.EndpointBuilder, "HEAD", path)
}
