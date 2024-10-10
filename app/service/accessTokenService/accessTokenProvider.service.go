package accessTokenService

import (
	"app/internal/bootstrap"
	"app/internal/common"
	libCommon "app/internal/lib/common"
	libError "app/internal/lib/error"
	accessTokenServicePort "app/port/accessToken"
	generalTokenServicePort "app/port/generalToken"
	jwtTokenServicePort "app/port/jwtTokenService"
	"app/repository"
	"app/service/noExpireTokenProvider"
	"app/valueObject"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	exp_duration   = time.Minute * 5
	claim_subject  = "sub"
	claim_audience = "aud"
	claim_expire   = "exp"

	//key_no_expire = "NO_EXPIRE_TOKEN"
)

var (
	ERR_INVALID_ACCESS_TOKEN_TYPE = errors.New("invalid access token format")
	ERR_INVALID_USER_CREDENTIALS  = errors.New("accessToken error: invalid username or password")
)

type (
	IAccessTokenHandler = accessTokenServicePort.IAccessTokenManipulator
	IGeneralToken       = generalTokenServicePort.IGeneralToken

	JWTAccessTokenManipulatorService struct {
		AudienceList               accessTokenServicePort.AccessTokenAudienceList
		JWTTokenManipulatorService jwtTokenServicePort.IAsymmetricJWTTokenManipulator
		UserRepo                   repository.IUser
		ExpDuration                time.Duration
		WithoutExpire              bool
		noExpireTokenProvider.NoExpireTokenProvider
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

func (this *JWTAccessTokenManipulatorService) GenerateBased(
	accessToken IAccessToken, ctx context.Context,
) (IAccessToken, error) {

	newAt, err := this.makeFor(accessToken.GetTenantUUID(), ctx)

	if err != nil {

		return nil, err
	}

	if v, ok := (accessToken.GetAuthData()).(*valueObject.AuthData); ok {

		newAt.customClaims.AuthData = v

	} else {

		return nil, libError.NewInternal(ERR_INVALID_TOKEN)
	}

	return newAt, nil
}

func (this *JWTAccessTokenManipulatorService) queryAndGenerate(
	condition bson.D, tenantUUID uuid.UUID, ctx context.Context,
) (*valueObject.AuthData, error) {

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
						{
							"pipeline", mongo.Pipeline{
								bson.D{
									{
										"$match", bson.D{
											{"tenantUUID", tenantUUID},
										},
									},
								},
							},
						},
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
						{
							"pipeline", mongo.Pipeline{
								bson.D{
									{
										"$match", bson.D{
											{"tenantUUID", tenantUUID},
										},
									},
								},
							},
						},
					},
				},
			},
			// bson.D{
			// 	{
			// 		"$set", bson.D{
			// 			{
			// 				"tenantAgentData", bson.D{
			// 					{
			// 						"$arrayElemAt", bson.A{"$tempTenantAgentDatas", 0},
			// 					},
			// 				},
			// 			},
			// 		},
			// 	},
			// },
			bson.D{
				{"$set",
					bson.D{
						{"isTenantAgent",
							bson.D{
								{"$cond",
									bson.D{
										{"if",
											bson.D{
												{"$ne",
													bson.A{
														"$tenantAgentData.0",
														primitive.Null{},
													},
												},
											},
										},
										{"then", true},
										{"else", false},
									},
								},
							},
						},
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
						{"isTenantAgent", 1},
						//{"tenantAgentData", 1},
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

		return nil, common.ERR_UNAUTHORIZED
	}

	authData := data[0]

	return authData, nil
}

func (this *JWTAccessTokenManipulatorService) makeFor(
	tenantUUID uuid.UUID, ctx context.Context,
) (*jwt_access_token, error) {

	token := this.JWTTokenManipulatorService.GenerateToken()

	customeClaims := &jwt_access_token_custom_claims{
		Issuer:     bootstrap.GetAppName(),
		IssuedAt:   jwt.NewNumericDate(time.Now()),
		TokenID:    "",
		TenantUUID: libCommon.PointerPrimitive(tenantUUID),
		AuthData:   nil,
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

	return "", libError.NewInternal(ERR_INVALID_ACCESS_TOKEN_TYPE)
}

func (this *JWTAccessTokenManipulatorService) DefaultExpireDuration() time.Duration {

	return exp_duration
}

func (this *JWTAccessTokenManipulatorService) GenerateFor(
	tenantUUID uuid.UUID, generalToken IGeneralToken, tokenID string, ctx context.Context,
) (accessTokenServicePort.IAccessToken, error) {

	authData, err := this.queryAndGenerate(
		bson.D{
			{"uuid", generalToken.GetUserUUID()},
		},
		tenantUUID,
		ctx,
	)
	fmt.Println(10, err)
	if err != nil {

		return nil, err
	}

	if authData == nil {

		return nil, errors.Join(common.ERR_UNAUTHORIZED)
	}

	at, err := this.makeFor(tenantUUID, ctx)
	fmt.Println(11, err)
	if err != nil {

		return nil, err
	}

	at.SetTokenID(tokenID)
	at.customClaims.AuthData = authData

	return at, nil
}
