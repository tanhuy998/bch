package boundedContext

import (
	authServicePort "app/port/auth"

	"github.com/kataras/iris/v12/hero"
)

type (
	AuthBoundedContext struct {
		authServicePort.IAddUserToCommandGroup
		authServicePort.ICheckCommandGroupUserRole
		authServicePort.ICheckUserInCommandGroup
		authServicePort.ICreateCommandGroup
		authServicePort.ICreateUser
		authServicePort.IGetAllRoles
		authServicePort.IGetCommandGroupUsers
		authServicePort.IGetParticipatedCommandGroups
		authServicePort.IGetSingleCommandGroup
		authServicePort.IGetSingleUser
		authServicePort.IGrantCommandGroupRolesToUser
		authServicePort.ILogIn
		authServicePort.IModifyUser
		authServicePort.IRefreshLogin
	}
)

func RegisterAuthBoundedContext(container *hero.Container) {

	container.Register(new(AuthBoundedContext)).Explicitly().EnableStructDependents()
}
