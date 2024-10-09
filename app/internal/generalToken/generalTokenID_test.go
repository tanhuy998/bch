package generalToken_test

import (
	"app/internal/generalToken"
	"encoding/json"
	"fmt"
	"testing"
)

func TestGeneralTokenID(t *testing.T) {

	type T struct {
		ID generalToken.GeneralTokenID `json:"id"`
	}

	o := &T{}

	b, err := json.Marshal(o)

	if err != nil {

		fmt.Println(err)
		return
	}

	fmt.Println(string(b))
}

func TestUnmarshal(t *testing.T) {

	type T struct {
		ID generalToken.GeneralTokenID `json:"id"`
	}

	o := &T{}

	err := json.Unmarshal([]byte(`{"id":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"}`), o)

	if err != nil {

		t.Error(err)
		return
	}

	temp := generalToken.GeneralTokenID{}

	fmt.Println(temp == o.ID)
}
