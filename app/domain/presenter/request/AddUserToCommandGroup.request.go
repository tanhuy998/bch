package requestPresenter

type (
	data_AddUserToCommandGroup struct {
		RoleUUIDs []string `json:"roleUUIDs"`
	}

	AddUserToCommandGroupRequest struct {
		GroupUUID string                     `param:"groupUUID" validate:"required"`
		UserUUID  string                     `param:"userUUID" validate:"required"`
		Data      data_AddUserToCommandGroup `json:"data"`
	}
)
