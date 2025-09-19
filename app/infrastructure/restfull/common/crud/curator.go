package crud

import "app/infrastructure/restfull/common/endpoint"

type (
	EndpointCurator struct {
		endpoint.APIEndpointCurator
	}
)

func (this *EndpointCurator) POST(path string) endpoint.IEndpointBuilder {

	return endpoint.Handle(&this.APIEndpointCurator, "POST", path)
}

func (this *EndpointCurator) DELETE(path string) endpoint.IEndpointBuilder {

	return endpoint.Handle(&this.APIEndpointCurator, "DELETE", path)
}

func (this *EndpointCurator) GET(path string) endpoint.IEndpointBuilder {

	return endpoint.Handle(&this.APIEndpointCurator, "GET", path)
}

func (this *EndpointCurator) HEAD(path string) endpoint.IEndpointBuilder {

	return endpoint.Handle(&this.APIEndpointCurator, "HEAD", path)
}

func (this *EndpointCurator) PUT(path string) endpoint.IEndpointBuilder {

	return endpoint.Handle(&this.APIEndpointCurator, "PUT", path)
}

func (this *EndpointCurator) PATCH(path string) endpoint.IEndpointBuilder {

	return endpoint.Handle(&this.APIEndpointCurator, "PATCH", path)
}
