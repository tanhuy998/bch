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
	ID_SPECTRUM_LIMIT      = 32 // byte
	SESSION_SPECTRUM_LIMIT = 8
)

type (
	GeneralTokenID          = generalToken.GeneralTokenID
	IRefreshTokenIDProvider = refreshtokenidServicePort.IRefreshTokenIDProvider

	RefreshTokenIDProviderService struct {
		UniqueIDGenerator uniqueIDServicePort.IUniqueIDGenerator
	}
)

func (this *RefreshTokenIDProviderService) Generate(generalTokenID GeneralTokenID) (string, error) {

	sessionID := make([]byte, SESSION_SPECTRUM_LIMIT)

	_, err := rand.Read(sessionID)

	if err != nil {

		return "", nil
	}

	return this.UniqueIDGenerator.Serve(append(generalTokenID[:], sessionID...))
}

func (this *RefreshTokenIDProviderService) Extract(
	ID string,
) (generalTokenID GeneralTokenID, sessionID [SESSION_SPECTRUM_LIMIT]byte, err error) {

	spectrum, err := this.UniqueIDGenerator.Decode(ID)

	if err != nil {

		err = libError.NewInternal(err)
		return
	}

	if len(spectrum) != ID_SPECTRUM_LIMIT {

		err = libError.NewInternal(fmt.Errorf("invalid refresh token id"))
		return
	}

	copy(generalTokenID[:], spectrum[:cap(generalTokenID)])
	copy(sessionID[:], spectrum[cap(generalTokenID):ID_SPECTRUM_LIMIT])

	return
}

// func (this *RefreshTokenIDProviderService) convertUUIDToBytes(u uuid.UUID) []byte {

// 	ret := make([]byte, 16)

// 	copy(ret, u[:])

// 	return ret
// }
