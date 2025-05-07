package getAssignmentGroupUnAssignedCommandGroupUsersDomain

import (
	libCommon "app/internal/lib/common"
	requestPresenter "app/presenter/request"
	"context"

	"github.com/google/uuid"
)

type (
	domain_input struct {
		requestPresenter.GetAssignmentGroupUnAssignedCommandGroupUsers
		executionCtx context.Context
	}
)

func mapDomainInput(presenter *requestPresenter.GetAssignmentGroupUnAssignedCommandGroupUsers) *domain_input {

	ret := &domain_input{
		*presenter,
		libCommon.Ternary[context.Context](
			presenter.GetAuthority().IsTenantAgent(),
			presenter.GetContext(),
			&non_tenant_agent_context{
				presenter.GetContext(),
			},
		),
	}

	return ret
}

func (this *domain_input) GetAssignmentGroupUUID() uuid.UUID {

	return *this.AssignmentGroupUUID
}

func (this *domain_input) GetContext() context.Context {

	return this.executionCtx
}
