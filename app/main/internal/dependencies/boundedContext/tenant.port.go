package boundedContext

import (
	irisIoc "app/internal/lib/iris/ioc"
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
		CreateTenantService         tenantServicePort.ICreateTenant
		CreatetenantAgentService    tenantServicePort.ICreateTenantAgent
		GetSingleTenantAgentService tenantServicePort.IGetSingleTenantAgent
		//tenantServicePort.IGetsingl
	}
)

func RegisterTenantBoundedContext(container *hero.Container) {

	irisIoc.BindDependency[tenantServicePort.IGetSingleTenant, getSingleTenantDomain.GetSingleTenantService](container, nil)
	irisIoc.BindDependency[tenantServicePort.IGetSingleTenantAgent, getSingleTenantAgentDomain.GetSingleTenantAgentService](container, nil)
	irisIoc.BindDependency[tenantServicePort.ICreateTenantAgent, createTenantAgentDomain.CreateTenantAgentService](container, nil)
	irisIoc.BindDependency[tenantServicePort.ICreateTenant, createTenantDomain.CreateTenantService](container, nil)

	irisIoc.BindDependency[
		usecasePort.IUseCase[requestPresenter.CreateTenantRequest, responsePresenter.CreateTenantResponse],
		createTenantDomain.CreateTenantUseCase,
	](container, nil)
	irisIoc.BindDependency[
		usecasePort.IUseCase[requestPresenter.GrantUserAsTenantAgent, responsePresenter.GrantUserAsTenantAgent],
		grantUserAsTenantAgentDomain.GrantUserAsTenantAgentUseCase,
	](container, nil)

	container.Register(new(TenantBoundedContext)).Explicitly().EnableStructDependents()
}
