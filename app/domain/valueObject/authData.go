package valueObject

import (
	accessTokenServicePort "app/adapter/accessToken"
	"app/domain/model"

	"github.com/google/uuid"
)

type (
	IParticipatedCommandGroup = accessTokenServicePort.IParticipatedCommandGroup

	UserParticipatedCommandGroup struct {
		CommandGroupUUID *uuid.UUID `json:"commandGroupUUID,omitempty" bson:"commandGroupUUID"`
		Role             string     `json:"role" bson:"role"`
	}

	AuthData struct {
		UserUUID                  *uuid.UUID                      `json:"-" bson:"uuid"`
		Name                      string                          `json:"name,omitempty" bson:"name"`
		TenantUUID                *uuid.UUID                      `json:"tenantUUID,omitempty" bson:"tenantUUID"`
		TenantAgentData           []*model.TenantAgent            `json:"tenantAgentData,omitempty" bson:"tenantAgentData"`
		ParticipatedCommandGroups []*UserParticipatedCommandGroup `json:"participatedGroups,omitempty" bson:"participatedGroups"`
	}
)

func (this *AuthData) GetTenantUUID() uuid.UUID {

	return *this.TenantUUID
}

func (this *AuthData) GetTenantAgentData() *model.TenantAgent {

	if len(this.TenantAgentData) == 0 {

		return nil
	}

	return this.TenantAgentData[0]
}

func (this *AuthData) GetParticipatedGroups() (ret []IParticipatedCommandGroup) {

	for _, val := range this.ParticipatedCommandGroups {

		ret = append(ret, val)
	}

	return ret
}

func (this *AuthData) IsTenantAgent() bool {

	return this.TenantAgentData != nil
}

func (this *AuthData) GetUserUUID() uuid.UUID {

	return *this.UserUUID
}

func (this *UserParticipatedCommandGroup) GetCommandGroupUUID() *uuid.UUID {

	return this.CommandGroupUUID
}
func (this *UserParticipatedCommandGroup) GetCommandGroupRoleName() string {

	return this.Role
}
