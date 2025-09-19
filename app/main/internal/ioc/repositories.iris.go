package ioc

import (
	repositoryConfig "app/infrastructure/http/common/config/repository"
	"app/infrastructure/http/common/config/repository/binding"
	"app/main/internal/dependencies/log"
	"app/main/internal/dependencies/repositories"

	"app/internal/db"
	"app/internal/db/driver/mongoDriver"
	"app/internal/db/query"
	"app/internal/db/relation"
	irisIoc "app/internal/lib/iris/ioc"
	iocOption "app/internal/lib/iris/ioc/option"

	"app/model"
	dbQueryTracerPort "app/port/dbQueryTracer"

	"app/repository"
	"app/service/mongoDBTracerService"

	"github.com/kataras/iris/v12/hero"
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

func InitializeDatabase(container *hero.Container /*app router.Party*/) {

	// var container *hero.Container = app.ConfigureContainer().EnableStructDependents().Container

	log.Main().Println("Initialize DBMS client...")
	client := db.GetClient()

	dbInstance := db.GetDB()

	container.Register(dbInstance).Explicitly()
	container.Register(client).Explicitly()

	irisIoc.BindDependency[repository.ITransactionDBClient, repository.MongoDBClient](container, nil)

	log.Main().Println("DBMS client initialized.")

	irisIoc.BindDependency[
		dbQueryTracerPort.IDBQueryTracer, mongoDBTracerService.DBQueryTracerService,
	](container, nil)

	/*
		Bind QueryBuilderGenerator as query.IQueryBuilderGenerator for agggregate
		to build complex query
	*/
	irisIoc.RegisterDependency(
		container, new(mongoDriver.ReadQueryBuilderGenerator),
		iocOption.StructDependents(true),
		iocOption.BindAs[query.IReadQueryBuilderGenerator](),
	)

	irisIoc.RegisterDependency(
		container, new(mongoDriver.ReadRelationQueryBuilderGenerator),
		iocOption.StructDependents(true),
		iocOption.BindAs[relation.IReadRelationQueryBuilderGenerator](),
	)

	log.Main().Println("Initialize Repositories...")

	repositoryConfig.BindRepositories(
		container,
		binding.ByRepositoryOf[model.Tenant](
			// new(repository.TenantRepository).Init(dbInstance),
			//new(mongoRepository.MongoCRUDRepository[model.Tenant]).Init(dbInstance, TENANT_COLLECTION_NAME),
			repositories.TenantRepository,
		),
		binding.ByRepositoryOf[model.TenantAgent](
			// new(repository.TenantAgentRepository).Init(dbInstance),
			// new(mongoRepository.MongoCRUDRepository[model.TenantAgent]).Init(dbInstance, TENANT_AGENT_COLLECTION_NAME),
			repositories.TenantAgentRepository,
		),
		binding.ByRepositoryOf[model.CommandGroup](
			// new(repository.CommandGroupRepository).Init(dbInstance),
			// new(mongoRepository.MongoCRUDRepository[model.CommandGroup]).Init(dbInstance, COMMAND_GROUP_COLLECTION_NAME),
			repositories.CommandGroupRepository,
		),
		binding.ByRepositoryOf[model.User](
			// new(repository.UserRepository).Init(dbInstance),
			//new(mongoRepository.MongoCRUDRepository[model.User]).Init(dbInstance, USER_COLLECTION_NAME),
			repositories.UserRepository,
		),
		binding.ByRepositoryOf[model.CommandGroupUser](
			// new(repository.CommandGroupUserRepository).Init(dbInstance),
			// new(mongoRepository.MongoCRUDRepository[model.CommandGroupUser]).Init(dbInstance, COMMAND_GROUP_USER_COLLECTION_NAME),
			repositories.CommandGroupUser,
		),
		binding.ByRepositoryOf[model.Assignment](
			// new(repository.AssignmentRepository).Init(dbInstance),
			// new(mongoRepository.MongoCRUDRepository[model.Assignment]).Init(dbInstance, ASSIGNMENT_COLLECTION_NAME),
			repositories.AssignmentRepository,
		),
		binding.ByRepositoryOf[model.AssignmentGroup](
			// new(repository.AssignmentGroupRepository).Init(dbInstance),
			// new(mongoRepository.MongoCRUDRepository[model.AssignmentGroup]).Init(dbInstance, ASSIGNMENT_GROUP_COLLECTION_NAME),
			repositories.AssignmentGroupRepository,
		),
		binding.ByRepositoryOf[model.AssignmentGroupMember](
			// new(repository.AssignmentGroupMemberRepository).Init(dbInstance),
			// new(mongoRepository.MongoCRUDRepository[model.AssignmentGroupMember]).Init(dbInstance, ASSIGNMENT_GROUP_MEMBER_COLLECTION_NAME),
			repositories.AssignmentGroupMemberRepository,
		),
		binding.ByRepositoryOf[model.UserSession](
			// new(repository.UserSessionRepository).Init(dbInstance),
			// new(mongoRepository.MongoCRUDRepository[model.UserSession]).Init(dbInstance, USER_SESSION_COLLECTION_NAME),
			repositories.UserSessionRepository,
		),
		binding.ByRepositoryOf[model.CommandGroupUserRole](
			// new(repository.CommandGroupUserRoleRepository).Init(dbInstance),
			// new(mongoRepository.MongoCRUDRepository[model.CommandGroupUserRole]).Init(dbInstance, COMMAND_GROUP_USER_ROLE_COLLECTION_NAME),
			repositories.CommandGroupUserRoleRepository,
		),
		binding.ByRepositoryOf[model.Role](
			// new(repository.RoleRepository).Init(dbInstance),
			// new(mongoRepository.MongoCRUDRepository[model.Role]).Init(dbInstance, ROLE_COLLECTION_NAME),
			repositories.RoleRepository,
		),
	)

	log.Main().Println("Repositories Initialized.")
}
