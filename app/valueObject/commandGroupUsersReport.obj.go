package valueObject

import (
	"app/model"

	"github.com/google/uuid"
)

type (
	CommandGroupUsersReport struct {
		GroupUUID uuid.UUID                      `json:"groupUUID"`
		GroupName string                         `json:"groupName"`
		Users     []*CommandGroupUsersReportData `json:"users"`
	}

	CommandGroupUsersReportData struct {
		UserUUID uuid.UUID     `json:"userUUID" bson:""`
		Name     string        `json:"name"`
		Username string        `json:"username"`
		Roles    []*model.Role `json:"roles"`
	}
)
