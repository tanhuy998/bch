package valueObject

import (
	"app/model"

	"github.com/google/uuid"
)

type (
	IParticipatedCommandGroup interface {
		GetCommandGroupUUID() *uuid.UUID
		GetCommandGroupRoleName() string
		HasRoles(name ...string) bool
	}

	IAuthorityData interface {
		GetUserUUID() uuid.UUID
		GetTenantUUID() uuid.UUID
		GetTenantAgentData() *model.TenantAgent
		GetParticipatedGroups() []IParticipatedCommandGroup
		IsTenantAgent() bool
	}

	// IParticipatedCommandGroup = accessTokenServicePort.IParticipatedCommandGroup

	UserParticipatedCommandGroup struct {
		CommandGroupUUID *uuid.UUID `json:"commandGroupUUID,omitempty" bson:"commandGroupUUID"`
		Roles            []string   `json:"roles" bson:"roles"`
	}

	/*
		Authority of a user to a specific tenant corresponding to a particular access token
	*/
	AuthData struct {
		ParticipatedCommandGroups []*UserParticipatedCommandGroup `json:"participatedCommandGroups,omitempty" bson:"participatedCommandGroups"`
		TenantAgentData           []model.TenantAgent             `json:"tenantAgentData,omitempty"`
		Name                      string                          `json:"name,omitempty" bson:"name,omitempty"`
		UserUUID                  *uuid.UUID                      `json:"uuid" bson:"uuid"`
		/*
			TenantUUId is the tenant which the user is belonged to
		*/
		TenantUUID *uuid.UUID `json:"tenantUUID,omitempty" bson:"tenantUUID"`

		/*
			IsAgent reports that the current user is agent of the corresponding tenant access token,
			not the user's tenant uuid.
			IsAgent is mostly used by Auth middlewware so it is unneccesary to use in domain level
		*/
		IsAgent bool `json:"isTenantAgent" bson:"isTenantAgent"`
	}
)

func (this *AuthData) GetTenantUUID() uuid.UUID {

	return *this.TenantUUID
}

func (this *AuthData) GetTenantAgentData() *model.TenantAgent {

	// if len(this.TenantAgentData) == 0 {

	// 	return nil
	// }

	// return this.TenantAgentData[0]

	return nil
}

func (this *AuthData) GetParticipatedGroups() []IParticipatedCommandGroup {

	ret := make([]IParticipatedCommandGroup, len(this.ParticipatedCommandGroups))

	for i, val := range this.ParticipatedCommandGroups {

		ret[i] = val
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

	return ""
}

func (this *UserParticipatedCommandGroup) HasRoles(names ...string) bool {

	if len(names) == 0 {

		return true
	}

	if len(this.Roles) < len(names) {

		return false
	}

	m := make(map[string]struct{})

	for _, v := range names {

		m[v] = struct{}{}
	}

	for _, rname := range this.Roles {

		if _, match := m[rname]; !match {

			continue
		}

		delete(m, rname)
	}

	return len(m) == 0
}
