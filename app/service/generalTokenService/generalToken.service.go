package generalTokenService

import (
	"app/internal/common"
	libCommon "app/internal/lib/common"
	libError "app/internal/lib/error"
	generalTokenServicePort "app/port/generalToken"
	generalTokenIDServicePort "app/port/generalTokenID"

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
	default_exp_duration = 15 * time.Minute
)

type (
	IGeneralToken           = generalTokenServicePort.IGeneralToken
	GeneralTokenManipulator struct {
		GeneralTokenIDProvider           generalTokenIDServicePort.IGeneralTokenIDProvider
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

	if val, ok := at.(*jwt_general_token); ok {

		return this.SymmetricTokenManipulatorService.SignString(val.jwt_token)
	}

	return "", libError.NewInternal(fmt.Errorf("TenantAccessTokenManipulatorService error: invalid tenant access token sign"))
}

func (this *GeneralTokenManipulator) makeFor(userUUID uuid.UUID, ctx context.Context) (IGeneralToken, error) {

	if userUUID == uuid.Nil {

		return nil, libError.NewInternal(fmt.Errorf("TenantAccessTokenManipulatorService error: nil tenant uuid given"))
	}

	token := this.SymmetricTokenManipulatorService.GenerateToken()

	generalToken, err := this.GeneralTokenIDProvider.Serve(userUUID)

	if err != nil {

		return nil, err
	}

	customClaims := &custom_claims{
		IssuedAt:       jwt.NewNumericDate(time.Now()),
		UserUUID:       libCommon.PointerPrimitive(userUUID),
		GeneralTokenID: libCommon.PointerPrimitive(generalToken),
		ExpireAt:       jwt.NewNumericDate(time.Now().Add(default_exp_duration)),
	}
	token.Claims = customClaims

	// if !this.IsNoExpire(ctx) {

	// 	customClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(default_exp_duration))
	// }

	return newFromToken(token)
}

func (this *GeneralTokenManipulator) GetDefaultExpireDuration() time.Duration {

	return default_exp_duration
}
