package refreshTokenIDService

import (
	refreshtokenidServicePort "app/adapter/refreshTokenidServicePort"
	uniqueIDServicePort "app/adapter/uniqueID"
	"crypto/rand"

	"github.com/google/uuid"
)

type (
	IRefreshTokenIDProvider = refreshtokenidServicePort.IRefreshTokenIDProvider

	RefreshTokenIDProviderService struct {
		UniqueIDGenerator uniqueIDServicePort.IUniqueIDGenerator
	}
)

func (this *RefreshTokenIDProviderService) Generate(userUUID uuid.UUID) (string, error) {

	sessionID := make([]byte, 8)

	_, err := rand.Read(sessionID)

	if err != nil {

		return "", nil
	}

	extraBytes := this.convertUUIDToBytes(userUUID)

	return this.UniqueIDGenerator.Serve(append(sessionID, extraBytes...))
}

func (this *RefreshTokenIDProviderService) convertUUIDToBytes(u uuid.UUID) []byte {

	ret := make([]byte, 16)

	copy(ret, u[0:15])

	return ret
}
