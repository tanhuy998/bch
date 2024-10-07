package generalTokenService

import (
	"app/internal/common"
	libCommon "app/internal/lib/common"
	libError "app/internal/lib/error"
	"app/port/generalTokenServicePort"
	jwtTokenServicePort "app/port/jwtTokenService"
	"app/service/noExpireTokenProvider"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	default_expire_duration = 2 * time.Hour
)

type (
	IGeneralToken           = generalTokenServicePort.IGeneralToken
	GeneralTokenManipulator struct {
		SymmetricTokenManipulatorService jwtTokenServicePort.ISymmetricJWTTokenManipulator
		noExpireTokenProvider.NoExpireTokenProvider
	}
)

func (this *GeneralTokenManipulator) Generate(
	userUUID uuid.UUID, ctx context.Context,
) (generalTokenServicePort.IGeneralToken, error) {

	at, err := this.makeFor(userUUID, ctx)

	if err != nil {

		return nil, err
	}

	return at, nil
}

func (this *GeneralTokenManipulator) Read(str string) (generalTokenServicePort.IGeneralToken, error) {

	custom_claims := new(custom_claims)

	token, err := this.SymmetricTokenManipulatorService.VerifyTokenStringCustomClaim(str, custom_claims)

	if errors.Is(err, common.ERR_INTERNAL) {

		return nil, err
	}

	if err != nil {

		return nil, libError.NewInternal(err)
	}

	return newFromToken(token)
}

func (this *GeneralTokenManipulator) SignString(at generalTokenServicePort.IGeneralToken) (string, error) {

	if val, ok := at.(*jwt_tenant_access_token); ok {

		return this.SymmetricTokenManipulatorService.SignString(val.jwt_token)
	}

	return "", libError.NewInternal(fmt.Errorf("TenantAccessTokenManipulatorService error: invalid tenant access token sign"))
}

func (this *GeneralTokenManipulator) makeFor(userUUID uuid.UUID, ctx context.Context) (ITenantAccessToken, error) {

	if userUUID == uuid.Nil {

		return nil, libError.NewInternal(fmt.Errorf("TenantAccessTokenManipulatorService error: nil tenant uuid given"))
	}

	token := this.SymmetricTokenManipulatorService.GenerateToken()

	customClaims := &custom_claims{
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
		UserUUID: libCommon.PointerPrimitive(userUUID),
	}
	token.Claims = customClaims

	if !this.IsNoExpire(ctx) {

		customClaims.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(default_expire_duration))
	}

	return newFromToken(token)
}

func (this *GeneralTokenManipulator) GetDefaultExpireDuration() time.Duration {

	return default_expire_duration
}
