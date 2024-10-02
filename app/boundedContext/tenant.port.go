package boundedContext

import (
	tenantServicePort "app/port/tenant"

	"github.com/kataras/iris/v12/hero"
)

type (
	TenantBoundedContext struct {
		tenantServicePort.ICreateTenant
		tenantServicePort.ICreateTenantAgent
		tenantServicePort.IGetSingleTenantAgent
	}
)

func RegisterTenantBoundedContext(container *hero.Container) {

	container.Register(new(TenantBoundedContext)).Explicitly().EnableStructDependents()
}
