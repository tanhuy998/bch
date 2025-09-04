package crud

import "app/infrastructure/restfull/common/endpoint"

type (
	/*
		Types that embed this class must not override BeforeActivation(mvc.BeforeActivation) method
	*/
	UpdateEndpointBuilder struct {
		endpoint.EndpointBuilder
	}
)

func (this *UpdateEndpointBuilder) PUT(path string) endpoint.IEndpointInitiator {

	return endpoint.Handle(&this.EndpointBuilder, "PUT", path)
}

func (this *UpdateEndpointBuilder) PATCH(path string) endpoint.IEndpointInitiator {

	return endpoint.Handle(&this.EndpointBuilder, "PATCH", path)
}
