package authService

import (
	libCommon "app/internal/lib/common"
	"app/model"
	"app/repository"
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// func Initialize(db *mongo.Database) {

// 	roleRepo := new(repository.RoleRepository).Init(db)

// 	InitializeRoles(roleRepo)
// }

func InitializeRoles(roleRepository repository.IRole) {

	init_entry_role(AUTH_COMMANDER_ROLE, roleRepository)
	init_entry_role(AUTH_MEMBER_ROLE, roleRepository)
}

func init_entry_role(roleName string, roleRepository repository.IRole) {

	res, err := roleRepository.Find(
		bson.D{
			{"name", roleName},
		},
		context.TODO(),
	)

	if err == mongo.ErrNoDocuments || res == nil {

		roleRepository.Create(
			&model.Role{
				UUID: libCommon.PointerPrimitive(uuid.New()),
				Name: roleName,
			},
			context.TODO(),
		)

		return
	}

	if err != nil {

		panic(err)
	}
}
