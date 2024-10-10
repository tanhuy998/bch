package valueObject

import (
	"app/model"
	accessTokenServicePort "app/port/accessToken"

	"github.com/google/uuid"
)

type (
	IParticipatedCommandGroup = accessTokenServicePort.IParticipatedCommandGroup

	UserParticipatedCommandGroup struct {
		CommandGroupUUID *uuid.UUID `json:"commandGroupUUID,omitempty" bson:"commandGroupUUID"`
		Role             string     `json:"role" bson:"role"`
	}

	AuthData struct {
		ParticipatedCommandGroups []*UserParticipatedCommandGroup `json:"participatedGroups,omitempty" bson:"participatedGroups"`
		Name                      string                          `json:"name,omitempty" bson:"name"`
		UserUUID                  *uuid.UUID                      `json:"uuid" bson:"uuid"`
		TenantUUID                *uuid.UUID                      `json:"tenantUUID,omitempty" bson:"tenantUUID"`
		TenantAgentData           *model.TenantAgent              `json:"tenantAgentData,omitempty"`
		IsAgent                   bool                            `json:"isTenantAgent" bson:"isTenantAgent"`
	}
)

func (this *AuthData) GetTenantUUID() uuid.UUID {

	return *this.TenantUUID
}

func (this *AuthData) GetTenantAgentData() *model.TenantAgent {

	return this.TenantAgentData
}

func (this *AuthData) GetParticipatedGroups() (ret []IParticipatedCommandGroup) {

	for _, val := range this.ParticipatedCommandGroups {

		ret = append(ret, val)
	}

	return ret
}

func (this *AuthData) IsTenantAgent() bool {

	return this.IsAgent
}

func (this *AuthData) GetUserUUID() uuid.UUID {

	if this.UserUUID == nil && this.GetTenantAgentData() != nil && this.GetTenantAgentData().UserUUID != nil {

		return *this.GetTenantAgentData().UserUUID
	}

	if this.UserUUID == nil {

		return uuid.Nil
	}

	return *this.UserUUID
}

func (this *UserParticipatedCommandGroup) GetCommandGroupUUID() *uuid.UUID {

	return this.CommandGroupUUID
}
func (this *UserParticipatedCommandGroup) GetCommandGroupRoleName() string {

	return this.Role
}
