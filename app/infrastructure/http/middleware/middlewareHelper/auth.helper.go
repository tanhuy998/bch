package middlewareHelper

import (
	accessTokenServicePort "app/port/accessToken"
	"time"
)

type (
	AccessToken         = accessTokenServicePort.IAccessToken
	AuthorityConstraint func(accessToken accessTokenServicePort.IAccessToken) bool
	// auth_err_body_reponse struct {
)

func beginCache() (markAsPass func(at AccessToken), markUnPass func(at AccessToken), lastMark func(at AccessToken) *bool) {

	cache := make(map[string]bool)

	markAsPass = func(at AccessToken) {

		key := at.GetTokenID()

		cache[key] = true

		if at.GetExpire() == nil {

			return
		}

		time.AfterFunc(
			time.Until(*at.GetExpire()),
			func() {

				delete(cache, key)
			},
		)
	}

	markUnPass = func(at AccessToken) {

		key := at.GetTokenID()

		cache[key] = false

		if at.GetExpire() == nil {

			return
		}

		time.AfterFunc(
			time.Until(*at.GetExpire()),
			func() {

				delete(cache, key)
			},
		)
	}

	lastMark = func(at AccessToken) *bool {

		key := at.GetTokenID()

		if state, ok := cache[key]; ok {

			return &state
		}

		return nil
	}

	return
}

func AuthRequireRoles(roleNames ...string) AuthorityConstraint {

	if len(roleNames) == 0 {

		panic("there is no role constraint")
	}

	m := make(map[string]struct{})

	for _, v := range roleNames {

		m[v] = struct{}{}
	}

	markAsPass, markUnPass, lastMark := beginCache()

	return func(accessToken accessTokenServicePort.IAccessToken) bool {

		switch state := lastMark(accessToken); {
		case accessToken == nil:
			return false
		case state != nil:
			return *state
		}

		a := accessToken.GetAuthData()

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

		pass := counter == len(m)

		if pass {

			markAsPass(accessToken)
		}

		markUnPass(accessToken)
		return pass
	}
}

func AuthRequireOneOfRoles(roleNames ...string) AuthorityConstraint {

	if len(roleNames) == 0 {

		panic("there is no role")
	}

	m := make(map[string]struct{})

	for _, v := range roleNames {

		m[v] = struct{}{}
	}

	markAsPass, markUnPass, lastMark := beginCache()

	return func(accessToken accessTokenServicePort.IAccessToken) bool {

		switch state := lastMark(accessToken); {
		case accessToken == nil:
			return false
		case state != nil:
			return *state
		}

		a := accessToken.GetAuthData()

		if a == nil {

			return false
		}

		for _, v := range a.GetParticipatedGroups() {

			if _, ok := m[v.GetCommandGroupRoleName()]; !ok {

				markAsPass(accessToken)
				return true
			}
		}

		markUnPass(accessToken)
		return false
	}
}

func AuthRequiredTenantAgentExceptOneOfRoles(roles ...string) AuthorityConstraint {

	f := AuthRequireOneOfRoles(roles...)

	return func(accessToken AccessToken) bool {

		switch {
		case accessToken == nil:
			return false
		case accessToken.GetAuthData() == nil:
			return false
		case accessToken.GetAuthData().IsTenantAgent():
			return true
		default:
			return f(accessToken)
		}
	}
}

func AuthRequiredTenantAgentExceptMeetRoles(roles ...string) AuthorityConstraint {

	f := AuthRequireRoles(roles...)

	return func(accessToken AccessToken) bool {

		switch {
		case accessToken == nil:
			return false
		case accessToken.GetAuthData() == nil:
			return false
		case accessToken.GetAuthData().IsTenantAgent():
			return true
		default:
			return f(accessToken)
		}
	}
}

func AuthRequireTenantAgent(accessToken accessTokenServicePort.IAccessToken) bool {

	if accessToken == nil {

		return false
	}

	a := accessToken.GetAuthData()

	if a == nil {

		return false
	}

	return a.IsTenantAgent()
}

func OneOfConstraints(c ...AuthorityConstraint) AuthorityConstraint {

	if len(c) == 0 {

		panic("there is no auth constraint")
	}

	return func(accessToken accessTokenServicePort.IAccessToken) bool {

		if accessToken == nil {

			return false
		}

		for _, f := range c {

			if f(accessToken) {

				return true
			}
		}

		return false
	}
}
