package jwtTokenService_test

import (
	jwtTokenService "app/service/jwtToken"
	"crypto/ecdsa"
	"fmt"
	"testing"

	"github.com/golang-jwt/jwt/v5"
)

const (
	example_token = `eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.e30.BuSfsENy900BJcETHud8iQq5ME0yxwEwdCf8vV3iPusMDObC9586HpWpjq_BK9uCVtSJ6LHThEKoMZcNCqWiOA`
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

	private_key *ecdsa.PrivateKey
	public_key  *ecdsa.PublicKey
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

func TestSigningTokenString(t *testing.T) {

	err := readKeys()

	if err != nil {

		t.Error(err)
		return
	}

	s := jwtTokenService.NewECDSAService(jwt.SigningMethodES256, *private_key, *public_key)

	token := s.GenerateToken()

	str, err := s.SignString(token)

	if err != nil {

		t.Error(err)
		return
	}

	fmt.Print(str)
}

func TestParseTokenFromString(t *testing.T) {

	err := readKeys()

	if err != nil {

		t.Error(err)
		return
	}

	s := jwtTokenService.NewECDSAService(jwt.SigningMethodES256, *private_key, *public_key)

	_, err = s.VerifyTokenString(example_token)

	if err != nil {

		t.Error(err)
		return
	}
}

func TestParseAndReadTokenFromString(t *testing.T) {

}
