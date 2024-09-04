package requestPresenter

type (
	CreateTenantAgentRequest struct {
		//Data *model.TenantAgent `json:"data" validate:"required"`
		Data *InputUser `json:"data" validate:"required"`
	}
)
