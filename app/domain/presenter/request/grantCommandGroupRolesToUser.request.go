package requestPresenter

type (
	GrantCommandGroupRolesToUserRequest struct {
		GroupUUID string   `param:"groupUUID"`
		UserUUID  string   `param:"userUUID"`
		Data      []string `json:"data"`
	}
)
