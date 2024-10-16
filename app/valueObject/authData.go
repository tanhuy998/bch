package valueObject

import (
	"app/model"

	"github.com/google/uuid"
)

type (
	IParticipatedCommandGroup interface {
		GetCommandGroupUUID() uuid.UUID
		GetCommandGroupRoleName() string
		HasRoles(name ...string) bool
	}

	IAuthorityData interface {
		GetUserUUID() uuid.UUID
		GetTenantUUID() uuid.UUID
		GetTenantAgentData() *model.TenantAgent
		GetParticipatedGroups() []IParticipatedCommandGroup
		JoinedGroups(groupUUIID ...uuid.UUID) bool
		QueryCommandGroup(commandGroupUUID uuid.UUID) ICommandGroupQueryBuilder
		//Joined(commandGroupUUID uuid.UUID) ICommandGroupQueryBuilder
		IsTenantAgent() bool
	}

	ICommandGroupQueryBuilder interface {
		HasRoles(name ...string) ICommandGroupQueryBuilder
		Done() bool
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
		P               []UserParticipatedCommandGroup `json:"-" bson:"participatedCommandGroups"`
		TenantAgentData []model.TenantAgent            `json:"tenantAgentData,omitempty"`
		Name            string                         `json:"name,omitempty" bson:"name,omitempty"`
		UserUUID        *uuid.UUID                     `json:"uuid" bson:"uuid"`
		/*
			TenantUUId is the tenant which the user is belonged to
		*/
		TenantUUID                *uuid.UUID             `json:"tenantUUID,omitempty" bson:"tenantUUID"`
		ParticipatedCommandGroups map[uuid.UUID][]string `json:"participatedCommandGroups,omitempty"`

		/*
			IsAgent reports that the current user is agent of the corresponding tenant access token,
			not the user's tenant uuid.
			IsAgent is mostly used by Auth middlewware so it is unneccesary to use in domain level
		*/
		IsAgent bool `json:"isTenantAgent" bson:"isTenantAgent"`
	}

	role_query struct {
		ref             map[uuid.UUID][]string
		lookupGroupUUID uuid.UUID
		lookupRoles     []string
	}
)

func (this *AuthData) Init() {

	this.ParticipatedCommandGroups = make(map[uuid.UUID][]string, 0)

	for _, v := range this.P {

		this.ParticipatedCommandGroups[*v.CommandGroupUUID] = v.Roles
	}

	this.P = nil
}

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

	// ret := make([]IParticipatedCommandGroup, len(this.P))

	// for i, val := range this.P {

	// 	ret[i] = &val
	// }

	ret := make([]IParticipatedCommandGroup, len(this.ParticipatedCommandGroups))

	i := 0

	for UUID, roles := range this.ParticipatedCommandGroups {

		ret[i] = &UserParticipatedCommandGroup{
			CommandGroupUUID: &UUID,
			Roles:            roles,
		}

		i++
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

func (this *AuthData) JoinedGroups(listUUID ...uuid.UUID) bool {

	if len(listUUID) == 0 && len(this.ParticipatedCommandGroups) == 0 {

		return true

	} else if len(listUUID) == 0 {

		return false
	}

	conter := 0

	for _, v := range listUUID {

		if _, ok := this.ParticipatedCommandGroups[v]; ok {

			conter++
		}
	}

	return conter == len(this.ParticipatedCommandGroups)
}

func (this *AuthData) QueryCommandGroup(commandGroupUUID uuid.UUID) ICommandGroupQueryBuilder {

	return &role_query{
		ref:             this.ParticipatedCommandGroups,
		lookupGroupUUID: commandGroupUUID,
	}
}

func (this *UserParticipatedCommandGroup) GetCommandGroupUUID() uuid.UUID {

	return *this.CommandGroupUUID
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

func (this *role_query) HasRoles(name ...string) ICommandGroupQueryBuilder {

	this.lookupRoles = name

	return this
}

func (this *role_query) Done() bool {

	if this.lookupGroupUUID == uuid.Nil {

		return false
	}

	refRoles, ok := this.ref[this.lookupGroupUUID]

	if !ok {

		return false
	}

	lookupMap := make(map[string]struct{}, 0)

	for _, v := range this.lookupRoles {

		lookupMap[v] = struct{}{}
	}

	for _, name := range refRoles {

		delete(lookupMap, name)
	}

	return len(lookupMap) == 0
}
