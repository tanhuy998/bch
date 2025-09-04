package crud

import "app/infrastructure/restfull/common/endpoint"

type (
	/*
		Types that embed this class must not override BeforeActivation(mvc.BeforeActivation) method
	*/
	UpdateEndpointCurator struct {
		endpoint.APIEndpointCurator
	}
)

func (this *UpdateEndpointCurator) PUT(path string) endpoint.IEndpointBuilder {

	return endpoint.Handle(&this.APIEndpointCurator, "PUT", path)
}

func (this *UpdateEndpointCurator) PATCH(path string) endpoint.IEndpointBuilder {

	return endpoint.Handle(&this.APIEndpointCurator, "PATCH", path)
}
