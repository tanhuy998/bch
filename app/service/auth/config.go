package authService

import (
	libCommon "app/internal/lib/common"
	"app/model"
	"app/repository"
	repositoryAPI "app/repository/api"
	"context"

	"github.com/google/uuid"
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

	// res, err := roleRepository.Filter(
	// 	// bson.D{
	// 	// 	{"name", roleName},
	// 	// },
	// 	func(filter repositoryAPI.IFilterGenerator) {

	// 		filter.Field("name").Equal(roleName)
	// 	},
	// ).Find(context.Background())

	// if err == mongo.ErrNoDocuments || res == nil {

	// 	roleRepository.Create(
	// 		&model.Role{
	// 			UUID: libCommon.PointerPrimitive(uuid.New()),
	// 			Name: roleName,
	// 		},
	// 		context.TODO(),
	// 	)

	// 	return
	// }

	err := roleRepository.Filter(
		func(filter repositoryAPI.IFilterGenerator) {

			filter.Field("name").Equal(roleName)
		},
	).Upsert(
		model.Role{
			UUID: libCommon.PointerPrimitive(uuid.New()),
			Name: roleName,
		},
		context.TODO(),
	)

	if err != nil {

		panic(err)
	}
}
