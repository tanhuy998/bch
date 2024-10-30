package jwtClaim

type (
	PrivateClaims struct {
		TokenType int `json:"typ"`
	}

	GenTokenPolicyClaim struct {
		Policies []GenTokenPolicyEnum `json:"pol"`
	}
)

type (
	GenTokenPolicyEnum int
)

const (
	POLICY_DEFAULT = GenTokenPolicyEnum(iota)
	POLICY_AT_NO_EXPIRE
)

const (
	GEN_TOKEN = iota
	REFRESH_TOKEN
	ACCESS_TOKEN
)

func SetupRefreshToken(obj *PrivateClaims) {

	obj.TokenType = REFRESH_TOKEN
}

func SetupAccessToken(obj *PrivateClaims) {

	obj.TokenType = ACCESS_TOKEN
}

func SetupGeneralToken(obj *PrivateClaims) {

	obj.TokenType = GEN_TOKEN
}
