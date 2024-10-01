package middlewareHelper

import (
	accessTokenServicePort "app/src/port/accessToken"
	"app/src/service/accessTokenService"
)

type (
	AuthorityConstraint func(authority accessTokenServicePort.IAccessTokenAuthData) bool
	// auth_err_body_reponse struct {
)

func AuthRequireRoles(roleNames ...string) AuthorityConstraint {

	m := make(map[string]struct{})

	for _, v := range roleNames {

		m[v] = struct{}{}
	}

	return func(a accessTokenServicePort.IAccessTokenAuthData) bool {

		if a == nil {

			return false
		}

		counter := 0

		for _, v := range a.GetParticipatedGroups() {

			if _, ok := m[v.GetCommandGroupRoleName()]; !ok {

				continue
			}

			counter++
		}

		return counter == len(m)
	}
}

func AuthRequireOneOfRoles(roleNames ...string) AuthorityConstraint {

	m := make(map[string]struct{})

	for _, v := range roleNames {

		m[v] = struct{}{}
	}

	return func(a accessTokenService.IAccessTokenAuthData) bool {

		if a == nil {

			return false
		}

		for _, v := range a.GetParticipatedGroups() {

			if _, ok := m[v.GetCommandGroupRoleName()]; !ok {

				return true
			}
		}

		return false
	}
}

func AuthRequireTenantAgent(a accessTokenService.IAccessTokenAuthData) bool {

	if a == nil {

		return false
	}

	return a.GetTenantAgentData() != nil
}
