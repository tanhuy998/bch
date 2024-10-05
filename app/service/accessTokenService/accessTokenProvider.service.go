package accessTokenService

import (
	"app/internal/bootstrap"
	"app/internal/common"
	"app/model"
	accessTokenServicePort "app/port/accessToken"
	jwtTokenServicePort "app/port/jwtTokenService"
	"app/repository"
	"app/valueObject"
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	exp_duration   = time.Minute * 5
	claim_subject  = "sub"
	claim_audience = "aud"
	claim_expire   = "exp"

	key_no_expire = "NO_EXPIRE_TOKEN"
)

var (
	ERR_INVALID_ACCESS_TOKEN_TYPE = errors.New("invalid access token format")
	ERR_INVALID_USER_CREDENTIALS  = errors.New("accessToken error: invalid username or password")
)

type (
	IAccessTokenHandler = accessTokenServicePort.IAccessTokenManipulator

	JWTAccessTokenManipulatorService struct {
		//JWtTokenSigningServive      jwtTokenServicePort.IJWTTokenSigning
		AudienceList accessTokenServicePort.AccessTokenAudienceList
		//FetchAuthDataService       authServiceAdapter.IFetchAuthData
		JWTTokenManipulatorService jwtTokenServicePort.IAsymmetricJWTTokenManipulator
		UserRepo                   repository.IUser
		ExpDuration                time.Duration
		WithoutExpire              bool
	}
)

func New(options ...AccessTokenManipulatorOption) *JWTAccessTokenManipulatorService {

	ret := new(JWTAccessTokenManipulatorService)

	for _, fn := range options {

		fn(ret)
	}

	return ret
}

func (this *JWTAccessTokenManipulatorService) Read(token_str string) (IAccessToken, error) {

	token, err := this.JWTTokenManipulatorService.VerifyTokenStringCustomClaim(token_str, &jwt_access_token_custom_claims{})

	if err != nil {

		return nil, err
	}

	return newFromToken(token)
}

func (this *JWTAccessTokenManipulatorService) GenerateByUserUUID(
	userUUID uuid.UUID, tokenID string, ctx context.Context,
) (at accessTokenServicePort.IAccessToken, err error) {

	at, err = this.queryAndGenerate(
		bson.D{
			{"uuid", userUUID},
		},
		ctx,
	)

	if err != nil || at == nil {

		return
	}

	at.SetTokenID(tokenID)
	return
}

func (this *JWTAccessTokenManipulatorService) GenerateByCredentials(
	model *model.User, tokenID string, ctx context.Context,
) (at accessTokenServicePort.IAccessToken, err error) {

	at, err = this.queryAndGenerate(
		bson.D{
			{"username", model.Username},
			//{"secret", model.Secret},
		},
		ctx,
	)

	if err != nil || at == nil {

		return
	}

	at.SetTokenID(tokenID)
	return
}

func (this *JWTAccessTokenManipulatorService) GenerateBased(
	accessToken IAccessToken, ctx context.Context,
) (IAccessToken, error) {

	newAt, err := this.makeFor(accessToken.GetUserUUID(), ctx)

	if err != nil {

		return nil, err
	}

	if v, ok := (accessToken.GetAuthData()).(*valueObject.AuthData); ok {

		newAt.customClaims.AuthData = v

	} else {

		return nil, ERR_INVALID_TOKEN
	}

	return newAt, nil
}

func (this *JWTAccessTokenManipulatorService) queryAndGenerate(
	condition bson.D, ctx context.Context,
) (accessTokenServicePort.IAccessToken, error) {

	data, err := repository.Aggregate[valueObject.AuthData](
		this.UserRepo.GetCollection(),
		mongo.Pipeline{
			bson.D{
				{
					"$match", condition,
				},
			},
			bson.D{
				{"$lookup",
					bson.D{
						{"from", "commandGroupUsers"},
						{"localField", "uuid"},
						{"foreignField", "userUUID"},
						{"as", "participatedCommandGroups"},
					},
				},
			},
			bson.D{
				{"$lookup",
					bson.D{
						{"from", "tenantAgents"},
						{"localField", "uuid"},
						{"foreignField", "userUUID"},
						{"as", "tenantAgentData"},
					},
				},
			},
			bson.D{
				{"$project",
					bson.D{
						{"uuid", 1},
						{"name", 1},
						{"username", 1},
						{"tenantUUID", 1},
						{"participatedCommandGroups", 1},
						{"tenantAgentData", 1},
					},
				},
			},
		},
		ctx,
	)

	if err != nil {

		return nil, errors.Join(
			common.ERR_INTERNAL,
			err,
		)
	}

	if len(data) == 0 {

		return nil, ERR_INVALID_USER_CREDENTIALS
	}

	authData := data[0]

	accessToken, err := this.makeFor(*authData.UserUUID, ctx)

	if err != nil {

		return nil, err
	}

	accessToken.customClaims.AuthData = authData

	return accessToken, nil
}

func (this *JWTAccessTokenManipulatorService) makeFor(userUUID uuid.UUID, ctx context.Context) (*jwt_access_token, error) {

	token := this.JWTTokenManipulatorService.GenerateToken()

	customeClaims := &jwt_access_token_custom_claims{
		jwt.RegisteredClaims{
			Subject:  userUUID.String(),
			Issuer:   bootstrap.GetAppName(),
			IssuedAt: jwt.NewNumericDate(time.Now()),
			//ExpiresAt: jwt.NewNumericDate(time.Now().Add(exp_duration)),
			//Audience:  jwt.ClaimStrings(bootstrap.GetHostNames()),
			//Issuer: ,
		},
		"",
		nil,
	}

	if !this.WithoutExpire && !this.IsNoExpire(ctx) {

		switch {
		case this.ExpDuration > 0:
			customeClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(this.ExpDuration))
		case this.ExpDuration <= 0:
			customeClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(exp_duration))
		}
	}

	token.Claims = customeClaims

	accesstoken, err := newFromToken(token)

	if err != nil {

		return nil, err
	}

	return accesstoken, nil
}

func (this *JWTAccessTokenManipulatorService) SignString(accessToken IAccessToken) (string, error) {

	if val, ok := accessToken.(*jwt_access_token); ok {

		return this.JWTTokenManipulatorService.SignString(val.jwt_token)
	}

	return "", ERR_INVALID_ACCESS_TOKEN_TYPE
}

func (this *JWTAccessTokenManipulatorService) DefaultExpireDuration() time.Duration {

	return exp_duration
}

func (this *JWTAccessTokenManipulatorService) CtxNoExpireKey() string {

	return key_no_expire
}

func (this *JWTAccessTokenManipulatorService) IsNoExpire(ctx context.Context) bool {

	v, ok := ctx.Value(key_no_expire).(bool)

	return ok && v
}
