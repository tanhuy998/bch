package crud

import "app/infrastructure/restfull/common/endpoint"

type (
	/*
		Types that embed this class must not override BeforeActivation(mvc.BeforeActivation) method
	*/
	DeleteEndpointBuilder struct {
		endpoint.EndpointBuilder
	}
)

func (this *DeleteEndpointBuilder) DELETE(path string) endpoint.IEndpointInitiator {

	return endpoint.Handle(&this.EndpointBuilder, "DELETE", path)
}
