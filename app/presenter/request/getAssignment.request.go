package requestPresenter

import (
	repositoryAPI "app/repository/api"
	"app/valueObject/requestInput"
	"time"
)

type (
	GetAssignments struct {
		// tenantUUID uuid.UUID
		// auth       accessTokenServicePort.IAccessTokenAuthData
		// ctx        context.Context
		requestInput.AuthorityInput
		requestInput.ContextInput
		requestInput.TenantMappingInput
		requestInput.RangePaginateInput
		requestInput.MongoCursorPaginateInput
		// PageNumber uint64             `url:"page"`
		// PageSize   uint64             `url:"size"`
		// Cursor  primitive.ObjectID `url:"p_cusor"`
		// IsPrev  bool               `url:p_prev`
		Expired bool `url:"expired"`
	}
)

func (this GetAssignments) GetExpiredFilter() bool {

	return this.Expired
}

func (this *GetAssignments) ApplyPaginateFilter(filter repositoryAPI.IFilterGenerator) {

	if this.Expired {

		filter.Field("deadline").LessThan(time.Now())
	}
}
