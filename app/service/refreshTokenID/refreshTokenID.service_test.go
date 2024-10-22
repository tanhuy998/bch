package refreshTokenIDService

import (
	"app/internal/generalToken"
	uniqueIDService "app/service/uniqueID"
	"testing"

	"github.com/google/uuid"
)

func TestGenerate(t *testing.T) {

	generalToken, err := generalToken.New(uuid.New())

	if err != nil {

		t.Errorf(err.Error())
	}

	idEngine, err := uniqueIDService.New(10)

	if err != nil {

		t.Errorf(err.Error())
	}

	obj := RefreshTokenIDProviderService{idEngine}

	refreshTokenID, err := obj.Generate(generalToken)

	if err != nil {

		t.Errorf(err.Error())
	}

	extracted, _, err := obj.Extract(refreshTokenID)

	if err != nil {

		t.Errorf(err.Error())
	}

	if extracted != generalToken {

		t.Errorf("failed")
	}
}
