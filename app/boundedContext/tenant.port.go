package boundedContext

import (
	libConfig "app/internal/lib/config"
	tenantServicePort "app/port/tenant"
	usecasePort "app/port/usecase"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"

	createTenantDomain "app/domain/tenant/createTenant"
	createTenantAgentDomain "app/domain/tenant/createTenantAgent"
	getSingleTenantAgentDomain "app/domain/tenant/getSingleTenantAgent"

	"github.com/kataras/iris/v12/hero"
)

type (
	TenantBoundedContext struct {
		tenantServicePort.ICreateTenant
		tenantServicePort.ICreateTenantAgent
		tenantServicePort.IGetSingleTenantAgent
		//tenantServicePort.IGetsingl
	}
)

func RegisterTenantBoundedContext(container *hero.Container) {

	libConfig.BindDependency[tenantServicePort.IGetSingleTenantAgent, getSingleTenantAgentDomain.GetSingleTenantAgentService](container, nil)
	libConfig.BindDependency[tenantServicePort.ICreateTenantAgent, createTenantAgentDomain.CreateTenantAgentService](container, nil)
	libConfig.BindDependency[tenantServicePort.ICreateTenant, createTenantDomain.CreateTenantService](container, nil)

	libConfig.BindDependency[
		usecasePort.IUseCase[requestPresenter.CreateTenantRequest, responsePresenter.CreateTenantResponse],
		createTenantDomain.CreateTenantUseCase,
	](container, nil)

	container.Register(new(TenantBoundedContext)).Explicitly().EnableStructDependents()
}
