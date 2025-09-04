package crud

import "app/infrastructure/restfull/common/endpoint"

type (
	/*
		Types that embed this class must not override BeforeActivation(mvc.BeforeActivation) method
	*/
	CreateEndpointBuilder struct {
		endpoint.EndpointBuilder
	}
)

func (this *CreateEndpointBuilder) POST(path string) endpoint.IEndpointInitiator {

	return endpoint.Handle(&this.EndpointBuilder, "POST", path)
}
