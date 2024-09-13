package refreshTokenService

import (
	accessTokenServicePort "app/adapter/accessToken"
	jwtTokenServicePort "app/adapter/jwtTokenService"
	refreshtokenServicePort "app/adapter/refreshToken"
	refreshTokenBlackListServicePort "app/adapter/refreshTokenBlackList"
	"app/domain/valueObject"
	refreshTokenIDService "app/service/refreshTokenID"
	"context"
	"errors"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	ERR_CANNOT_REVOKE_REFRESH_TOKEN = errors.New("RefreshTokenManipulatorService error: cannot revoke refresh token")
)

type (
	IRefreshTokenHandler = refreshtokenServicePort.IRefreshTokenHadler
	IRefreshToken        = refreshtokenServicePort.IRefreshToken
	RefreshTokenID       = refreshtokenServicePort.RefreshTokenID

	RefreshTokenManipulatorService struct {
		AudienceList           accessTokenServicePort.AccessTokenAudienceList
		RefreshTokenIDProvider refreshTokenIDService.IRefreshTokenIDProvider
		JWTTokenService        jwtTokenServicePort.ISymmetricJWTTokenManipulator
		RefreshTokenBlackList  refreshTokenBlackListServicePort.IRefreshTokenBlackListManipulator[string, refreshTokenBlackListServicePort.IRefreshTokenBlackListPayload]
	}
)

func (this *RefreshTokenManipulatorService) Generate(userUUID uuid.UUID, ctx context.Context) (IRefreshToken, error) {

	token := this.JWTTokenService.GenerateToken()

	refreshTokenID, err := this.RefreshTokenIDProvider.Generate(userUUID)

	if err != nil {

		return nil, err
	}

	token.Claims = jwt_refresh_token_custom_claims{
		jwt.RegisteredClaims{},
		refreshTokenID,
	}

	return newFromToken(token)
}
func (this *RefreshTokenManipulatorService) Revoke(RefreshTokenID string) error {

	payload := &valueObject.RefreshTokenBlackListPayload{}

	success, err := this.RefreshTokenBlackList.Set(RefreshTokenID, payload)

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
