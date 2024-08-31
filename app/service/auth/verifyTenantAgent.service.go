package authService

import "github.com/google/uuid"

type (
	IVerifyTenantAgent interface {
		Serve(agentUUID uuid.UUID) error
	}

	VerifyTenantAgentService struct {
	}
)

func (this *VerifyTenantAgentService) Serve(agentUUID uuid.UUID) error {

	return nil
}
