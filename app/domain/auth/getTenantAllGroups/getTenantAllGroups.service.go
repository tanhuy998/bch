package getTenantAllGroupsDomain

import (
	"app/model"
	authServicePort "app/port/auth"
	"app/repository"
	paginateUseCase "app/unitOfWork/genericUsecase/paginate"
)

type (
	GetTenantAllGroupService struct {
		//CommandGroupRepo repository.ICommandGroup
		paginateUseCase.PaginateUseCase[
			repository.ICommandGroup,
			model.CommandGroup,
			interface{},
		]
	}
)

func (this *GetTenantAllGroupService) Serve(input authServicePort.GetTenantAllGroupsInput) ([]model.CommandGroup, error) {

	// return repository.Aggregate[model.CommandGroup](
	// 	this.CommandGroupRepo.GetCollection(),
	// 	mongo.Pipeline{
	// 		bson.D{
	// 			{
	// 				"$match", bson.D{
	// 					{"tenantUUID", tenantUUID},
	// 				},
	// 			},
	// 		},
	// 	},
	// 	ctx,
	// )

	// return this.CommandGroupRepo.Filter(
	// 	func(filter repositoryAPI.IFilterGenerator) {
	// 		filter.Field("tenantUUID").Equal(input.GetTenantUUID())
	// 	},
	// ).Find(ctx)

	return this.PaginateUseCase.UseCustomPaginator(
		input.GetTenantUUID(), input.GetGeneralPaginator(), input.GetContext(),
	)
}
