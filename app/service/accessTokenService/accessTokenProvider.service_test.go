package accessTokenService

import (
	jwtTokenService "app/service/jwtToken"
	"crypto/ecdsa"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	private_pem = []byte(`-----BEGIN EC PRIVATE KEY-----
MHcCAQEEIOvT8A6VMvsJNoK5Usln8xvveyu3waaTkcmn1XrckWZ1oAoGCCqGSM49
AwEHoUQDQgAEFz9WMsJtxT+HmLqyLyF1dVGrPQ53LhtF8LyleJsE6vk4JmXOHn6O
pwYfzrJVKdp7o12CnQOLJPAItxtvCr20IQ==
-----END EC PRIVATE KEY-----`)
	public_pem = []byte(`-----BEGIN PUBLIC KEY-----
MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEFz9WMsJtxT+HmLqyLyF1dVGrPQ53
LhtF8LyleJsE6vk4JmXOHn6OpwYfzrJVKdp7o12CnQOLJPAItxtvCr20IQ==
-----END PUBLIC KEY-----`)

	private_key  *ecdsa.PrivateKey
	public_key   *ecdsa.PublicKey
	userUUID_str = `9f04a34b-c101-4f30-8f71-46c130ab9d76`
)

func readKeys() error {

	priv, err := jwt.ParseECPrivateKeyFromPEM(private_pem)

	if err != nil {

		return err
	}

	private_key = priv

	pub, err := jwt.ParseECPublicKeyFromPEM(public_pem)

	public_key = pub

	return err
}

func TestGenerateAccessToken(t *testing.T) {

	err := readKeys()

	if err != nil {

		t.Error(err)
		return
	}

	s := &JWTAccessTokenManipulatorService{
		JWTTokenManipulatorService: jwtTokenService.NewECDSAService(jwt.SigningMethodES256, *private_key, *public_key),
	}

	userUUID, err := uuid.Parse(userUUID_str)

	if err != nil {

		t.Error(err)
		return
	}

	at, err := s.makeFor(userUUID)

	if err != nil {

		t.Error(err)
		return
	}

	if *at.GetUserUUID() != userUUID {

		t.Error("fail")
		return
	}
}

func TestVerifyAccessTokenFromString(t *testing.T) {

	err := readKeys()

	if err != nil {

		t.Error(err)
		return
	}

	s := &JWTAccessTokenManipulatorService{
		JWTTokenManipulatorService: jwtTokenService.NewECDSAService(jwt.SigningMethodES256, *private_key, *public_key),
	}

	userUUID, err := uuid.Parse(userUUID_str)

	if err != nil {

		t.Error(err)
		return
	}

	a, err := s.makeFor(userUUID)

	if err != nil {

		t.Error(err)
		return
	}

	at_str, err := s.SignString(a)

	if err != nil {

		t.Error(err)
		return
	}

	b, err := s.Read(at_str)

	if err != nil {

		t.Error(err)
		return
	}

	if b.GetUserUUID().String() != userUUID_str {

		t.Error("fail")
		return
	}
}
