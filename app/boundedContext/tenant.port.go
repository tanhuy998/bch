package boundedContext

import (
	libConfig "app/internal/lib/config"
	tenantServicePort "app/port/tenant"
	usecasePort "app/port/usecase"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"

	createTenantDomain "app/domain/tenant/createTenant"
	createTenantAgentDomain "app/domain/tenant/createTenantAgent"
	getSingleTenantDomain "app/domain/tenant/getSingleTenant"
	getSingleTenantAgentDomain "app/domain/tenant/getSingleTenantAgent"
	grantUserAsTenantAgentDomain "app/domain/tenant/grantUserAsTenantAgent"

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

	libConfig.BindDependency[tenantServicePort.IGetSingleTenant, getSingleTenantDomain.GetSingleTenantService](container, nil)
	libConfig.BindDependency[tenantServicePort.IGetSingleTenantAgent, getSingleTenantAgentDomain.GetSingleTenantAgentService](container, nil)
	libConfig.BindDependency[tenantServicePort.ICreateTenantAgent, createTenantAgentDomain.CreateTenantAgentService](container, nil)
	libConfig.BindDependency[tenantServicePort.ICreateTenant, createTenantDomain.CreateTenantService](container, nil)

	libConfig.BindDependency[
		usecasePort.IUseCase[requestPresenter.CreateTenantRequest, responsePresenter.CreateTenantResponse],
		createTenantDomain.CreateTenantUseCase,
	](container, nil)
	libConfig.BindDependency[
		usecasePort.IUseCase[requestPresenter.GrantUserAsTenantAgent, responsePresenter.GrantUserAsTenantAgent],
		grantUserAsTenantAgentDomain.GrantUserAsTenantAgentUseCase,
	](container, nil)

	container.Register(new(TenantBoundedContext)).Explicitly().EnableStructDependents()
}
