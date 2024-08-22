package authService

import (
	"app/domain/model"
	"app/repository"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type (
	IGetAllRoles interface {
		Serve() ([]*model.Role, error)
	}

	GetAllRolesService struct {
		RoleRepo repository.IRole
	}
)

func (this *GetAllRolesService) Serve() ([]*model.Role, error) {

	return this.RoleRepo.FindMany(bson.D{{}}, context.TODO())
}
