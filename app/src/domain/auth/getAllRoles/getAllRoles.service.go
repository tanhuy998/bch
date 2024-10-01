package getAllRoleDomain

import (
	"app/src/model"
	"app/src/repository"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type (
	GetAllRolesService struct {
		RoleRepo repository.IRole
	}
)

func (this *GetAllRolesService) Serve() ([]*model.Role, error) {

	return this.RoleRepo.FindMany(bson.D{{}}, context.TODO())
}
