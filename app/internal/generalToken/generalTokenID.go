package generalToken

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"github.com/google/uuid"
)

var (
	ErrInvalidTextGeneralTokenID = fmt.Errorf("invalid GeneralTokenID length")
)

var (
	Nil = [24]byte{}
)

type (
	GeneralTokenID [24]byte
)

func New(userUUId uuid.UUID) (GeneralTokenID, error) {

	var ret GeneralTokenID

	baseID := make([]byte, 8)

	_, err := rand.Read(baseID)

	if err != nil {

		return ret, err
	}

	copy(ret[:8], baseID)
	copy(ret[8:], userUUId[:])

	return ret, nil
}

func (g GeneralTokenID) String() string {
	return base64.StdEncoding.EncodeToString((g)[:])
}

func (g GeneralTokenID) MarshalText() (text []byte, err error) {

	return base64.StdEncoding.AppendEncode([]byte{}, g[:]), nil
}

func (g *GeneralTokenID) UnmarshalText(text []byte) error {

	b := make([]byte, 0)

	b, err := base64.StdEncoding.AppendDecode(b, text)

	if err != nil {

		return err
	}

	if len(b) != 24 {

		return ErrInvalidTextGeneralTokenID
	}

	copy(g[:], b)
	return nil
}

func (g GeneralTokenID) GetBaseID() [8]byte {

	var ret [8]byte
	copy(ret[:], g[:8])
	return ret
}

func (g GeneralTokenID) GetUserUUID() uuid.UUID {

	var ret uuid.UUID
	copy(ret[:], g[8:])
	return ret
}

func (g GeneralTokenID) MarshalBinary() (data []byte, err error) {

	return g[:], nil
}

func (g *GeneralTokenID) UnmarshalBinary(data []byte) error {

	if len(data) != 24 {

		return fmt.Errorf("invalid Unmarshalled value (expected 24 bytes but got %d bytes)", len(data))
	}

	copy(g[:], data)
	return nil
}
