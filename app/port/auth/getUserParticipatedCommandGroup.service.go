package authServicePort

import (
	"app/model"
	paginateServicePort "app/port/paginate"
	"app/unitOfWork/aggregate"
	"context"

	"github.com/google/uuid"
)

type (
	GetUserParticipatedCommandGroupInput interface {
		aggregate.IDomainContext
		paginateServicePort.IGeneralPaginator
		GetRequestedUserUUID() uuid.UUID
		GetContext() context.Context
	}

	IGetUserParticipatedCommandGroups[Entity_T any] interface {
		Serve(
			//tenantUUID uuid.UUID, userUUID uuid.UUID, ctx context.Context,
			input GetUserParticipatedCommandGroupInput,
		) ([]Entity_T, error)
		SearchAndRetrieveByModel(
			searchModel *model.User, ctx context.Context,
		) ([]Entity_T, error)
	}
)
