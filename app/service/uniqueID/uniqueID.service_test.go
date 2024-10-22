package uniqueIDService

import (
	"crypto/rand"
	"fmt"
	"testing"

	"github.com/google/uuid"
)

func TestGenerateOddInput(t *testing.T) {

	obj, err := New(10)

	if err != nil {

		t.Errorf(err.Error())
		return
	}

	extraBytes := [5]byte{}
	_, err = rand.Read(extraBytes[:])

	if err != nil {

		t.Errorf(err.Error())
		return
	}

	id := uuid.New()

	buffer := make([]byte, len(id)+len(extraBytes))

	copy(buffer[:len(id)], id[:])
	copy(buffer[len(id):], extraBytes[:])

	str, err := obj.Serve(buffer)

	if err != nil {

		t.Errorf(err.Error())
		return
	}

	extractedBytes, err := obj.Decode(str)

	if err != nil {

		t.Errorf(err.Error())
		return
	}

	comparedBuffer := [16 + 5]byte{}
	extractedBuffer := [16 + 5]byte{}

	copy(comparedBuffer[:], buffer)
	copy(extractedBuffer[:], extractedBytes)

	if extractedBuffer != comparedBuffer {
		fmt.Println(buffer, extractedBuffer, comparedBuffer)
		t.Errorf("failed")
	}
}
