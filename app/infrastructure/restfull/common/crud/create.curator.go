package crud

import "app/infrastructure/restfull/common/endpoint"

type (
	/*
		Types that embed this class must not override BeforeActivation(mvc.BeforeActivation) method
	*/
	CreateEndpointCurator struct {
		endpoint.APIEndpointCurator
	}
)

func (this *CreateEndpointCurator) POST(path string) endpoint.IEndpointBuilder {

	return endpoint.Handle(&this.APIEndpointCurator, "POST", path)
}
