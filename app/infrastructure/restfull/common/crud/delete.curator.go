package crud

import "app/infrastructure/restfull/common/endpoint"

type (
	/*
		Types that embed this class must not override BeforeActivation(mvc.BeforeActivation) method
	*/
	DeleteEndpointCurator struct {
		endpoint.APIEndpointCurator
	}
)

func (this *DeleteEndpointCurator) DELETE(path string) endpoint.IEndpointBuilder {

	return endpoint.Handle(&this.APIEndpointCurator, "DELETE", path)
}
