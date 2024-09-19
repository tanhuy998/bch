package refreshTokenService

import (
	accessTokenServicePort "app/adapter/accessToken"
	jwtTokenServicePort "app/adapter/jwtTokenService"
	refreshTokenServicePort "app/adapter/refreshToken"
	refreshTokenBlackListServicePort "app/adapter/refreshTokenBlackList"
	"app/domain/valueObject"
	refreshTokenIDService "app/service/refreshTokenID"
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	exp_duration = time.Hour * 2
)

var (
	ERR_CANNOT_REVOKE_REFRESH_TOKEN = errors.New("RefreshTokenManipulatorService error: cannot revoke refresh token")
)

type (
	IRefreshTokenManipulator = refreshTokenServicePort.IRefreshTokenManipulator
	IRefreshToken            = refreshTokenServicePort.IRefreshToken

	RefreshTokenManipulatorService struct {
		AudienceList           accessTokenServicePort.AccessTokenAudienceList
		RefreshTokenIDProvider refreshTokenIDService.IRefreshTokenIDProvider
		JWTTokenService        jwtTokenServicePort.ISymmetricJWTTokenManipulator
		RefreshTokenBlackList  refreshTokenBlackListServicePort.IRefreshTokenBlackListManipulator
	}
)

func (this *RefreshTokenManipulatorService) Generate(userUUID uuid.UUID, ctx context.Context) (IRefreshToken, error) {

	token := this.JWTTokenService.GenerateToken()

	refreshTokenID, err := this.RefreshTokenIDProvider.Generate(userUUID)

	if err != nil {

		return nil, err
	}

	token.Claims = &jwt_refresh_token_custom_claims{
		jwt.RegisteredClaims{
			Subject:   userUUID.String(),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(exp_duration)),
		},
		refreshTokenID,
	}

	return newFromToken(token)
}
func (this *RefreshTokenManipulatorService) Revoke(RefreshTokenID string, ctx context.Context) error {

	payload := &valueObject.RefreshTokenBlackListPayload{}

	success, err := this.RefreshTokenBlackList.Set(RefreshTokenID, payload, ctx)

	if err != nil {

		return err
	}

	if !success {

		return ERR_CANNOT_REVOKE_REFRESH_TOKEN
	}

	return nil
}

func (this *RefreshTokenManipulatorService) SignString(refreshToken IRefreshToken) (string, error) {

	if v, ok := any(refreshToken).(*jwt_refresh_token); ok {

		return this.JWTTokenService.SignString(v.jwt_token)
	}

	return "", ERR_INVALID_TOKEN
}

func (this *RefreshTokenManipulatorService) Read(str string) (IRefreshToken, error) {

	token, err := this.JWTTokenService.VerifyTokenString(str)

	if err != nil {

		return nil, err
	}

	return newFromToken(token)
}

func (this *RefreshTokenManipulatorService) DefaultExpireDuration() time.Duration {

	return exp_duration
}
