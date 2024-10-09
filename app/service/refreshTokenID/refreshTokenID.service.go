package refreshTokenIDService

import (
	"app/internal/generalToken"
	libError "app/internal/lib/error"
	refreshtokenidServicePort "app/port/refreshTokenID"
	uniqueIDServicePort "app/port/uniqueID"
	"crypto/rand"
	"fmt"
)

const (
	ID_SPECTRUM_SIZE = 32
)

type (
	GeneralTokenID          = generalToken.GeneralTokenID
	IRefreshTokenIDProvider = refreshtokenidServicePort.IRefreshTokenIDProvider

	RefreshTokenIDProviderService struct {
		UniqueIDGenerator uniqueIDServicePort.IUniqueIDGenerator
	}
)

func (this *RefreshTokenIDProviderService) Generate(generalTokenID GeneralTokenID) (string, error) {

	sessionID := make([]byte, 8)

	_, err := rand.Read(sessionID)

	if err != nil {

		return "", nil
	}

	extraBytes := make([]byte, len(generalTokenID))

	return this.UniqueIDGenerator.Serve(append(sessionID, extraBytes...))
}

func (this *RefreshTokenIDProviderService) Extract(
	ID string,
) (generalTokenID GeneralTokenID, sessionID [8]byte, err error) {

	spectrum, err := this.UniqueIDGenerator.Decode(ID)

	if err != nil {

		err = libError.NewInternal(err)
		return
	}

	if len(spectrum) != ID_SPECTRUM_SIZE {

		err = libError.NewInternal(fmt.Errorf("invalid refresh token id"))
		return
	}

	copy(generalTokenID[:], spectrum[:24])
	copy(sessionID[:], spectrum[24:ID_SPECTRUM_SIZE])

	return
}

// func (this *RefreshTokenIDProviderService) convertUUIDToBytes(u uuid.UUID) []byte {

// 	ret := make([]byte, 16)

// 	copy(ret, u[:])

// 	return ret
// }
