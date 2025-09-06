package repositories

import (
	"app/internal/db"
	"app/model"
	mongoRepository "app/repository/driver/mongod"
)

const (
	CAMPAIGN_COLLETCTION = "campaigns"
	CANDIATE_COLLECTIONN = "candidates"
)

const (
	TENANT_COLLECTION_NAME                  = "tenants"
	TENANT_AGENT_COLLECTION_NAME            = "tenantAgents"
	COMMAND_GROUP_COLLECTION_NAME           = "commandGroups"
	COMMAND_GROUP_USER_COLLECTION_NAME      = "commandGroupUsers"
	COMMAND_GROUP_USER_ROLE_COLLECTION_NAME = "commandGroupUserRoles"
	USER_COLLECTION_NAME                    = "users"
	ASSIGNMENT_COLLECTION_NAME              = "assignments"
	ASSIGNMENT_GROUP_COLLECTION_NAME        = "assignmentGroups"
	ASSIGNMENT_GROUP_MEMBER_COLLECTION_NAME = "assignmentGroupMembers"
	USER_SESSION_COLLECTION_NAME            = "userSessions"
	ROLE_COLLECTION_NAME                    = "roles"
)

// func InitDomainIndexes(db *mongo.Database) {

// 	db.Collection(CAMPAIGN_COLLETCTION).Indexes().CreateMany(
// 		context.TODO(),
// 		[]mongo.IndexModel{
// 			mongo.IndexModel{
// 				Keys: "uuid",
// 				Options: &options.IndexOptions{
// 					Name:   libCommon.PointerPrimitive("uuid"),
// 					Unique: libCommon.PointerPrimitive(true),
// 				},
// 			},
// 		},
// 	)

// 	db.Collection(CANDIATE_COLLECTIONN).Indexes().CreateMany(
// 		context.TODO(),
// 		[]mongo.IndexModel{
// 			mongo.IndexModel{
// 				Keys: "uuid",
// 				Options: &options.IndexOptions{
// 					Name:   libCommon.PointerPrimitive("uuid"),
// 					Unique: libCommon.PointerPrimitive(true),
// 				},
// 			},
// 			mongo.IndexModel{
// 				Keys: "idNumber",
// 				Options: &options.IndexOptions{
// 					Name:   libCommon.PointerPrimitive("idNumber"),
// 					Unique: libCommon.PointerPrimitive(true),
// 				},
// 			},
// 		},
// 	)

// 	initProductionIndexes(db)

// }

// func initProductionIndexes(db *mongo.Database) {

// }

var (
	TenantRepository                *mongoRepository.MongoCRUDRepository[model.Tenant]
	TenantAgentRepository           *mongoRepository.MongoCRUDRepository[model.TenantAgent]
	CommandGroupRepository          *mongoRepository.MongoCRUDRepository[model.CommandGroup]
	CommandGroupUser                *mongoRepository.MongoCRUDRepository[model.CommandGroupUser]
	CommandGroupUserRoleRepository  *mongoRepository.MongoCRUDRepository[model.CommandGroupUserRole]
	UserRepository                  *mongoRepository.MongoCRUDRepository[model.User]
	AssignmentRepository            *mongoRepository.MongoCRUDRepository[model.Assignment]
	AssignmentGroupRepository       *mongoRepository.MongoCRUDRepository[model.AssignmentGroup]
	AssignmentGroupMemberRepository *mongoRepository.MongoCRUDRepository[model.AssignmentGroupMember]
	UserSessionRepository           *mongoRepository.MongoCRUDRepository[model.UserSession]
	RoleRepository                  *mongoRepository.MongoCRUDRepository[model.Role]
)

func init() {

	TenantRepository = new(mongoRepository.MongoCRUDRepository[model.Tenant]).Init(db.GetDB(), TENANT_COLLECTION_NAME)
	TenantAgentRepository = new(mongoRepository.MongoCRUDRepository[model.TenantAgent]).Init(db.GetDB(), TENANT_AGENT_COLLECTION_NAME)
	CommandGroupRepository = new(mongoRepository.MongoCRUDRepository[model.CommandGroup]).Init(db.GetDB(), COMMAND_GROUP_COLLECTION_NAME)
	CommandGroupUser = new(mongoRepository.MongoCRUDRepository[model.CommandGroupUser]).Init(db.GetDB(), COMMAND_GROUP_USER_COLLECTION_NAME)
	CommandGroupUserRoleRepository = new(mongoRepository.MongoCRUDRepository[model.CommandGroupUserRole]).Init(db.GetDB(), COMMAND_GROUP_USER_ROLE_COLLECTION_NAME)
	UserRepository = new(mongoRepository.MongoCRUDRepository[model.User]).Init(db.GetDB(), USER_COLLECTION_NAME)
	AssignmentRepository = new(mongoRepository.MongoCRUDRepository[model.Assignment]).Init(db.GetDB(), ASSIGNMENT_COLLECTION_NAME)
	AssignmentGroupRepository = new(mongoRepository.MongoCRUDRepository[model.AssignmentGroup]).Init(db.GetDB(), ASSIGNMENT_GROUP_COLLECTION_NAME)
	AssignmentGroupMemberRepository = new(mongoRepository.MongoCRUDRepository[model.AssignmentGroupMember]).Init(db.GetDB(), ASSIGNMENT_GROUP_MEMBER_COLLECTION_NAME)
	UserSessionRepository = new(mongoRepository.MongoCRUDRepository[model.UserSession]).Init(db.GetDB(), USER_SESSION_COLLECTION_NAME)
	RoleRepository = new(mongoRepository.MongoCRUDRepository[model.Role]).Init(db.GetDB(), ROLE_COLLECTION_NAME)
}
