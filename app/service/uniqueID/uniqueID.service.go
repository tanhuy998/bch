package uniqueIDService

import (
	libError "app/internal/lib/error"
	uniqueIDServicePort "app/port/uniqueID"
	"encoding/binary"

	"github.com/sqids/sqids-go"
)

const (
	ENGINE_UNIT_SIZE = 8 // byte length
)

type (
	IUniqueIDGenerator = uniqueIDServicePort.IUniqueIDGenerator

	uniqueIDService struct {
		engine *sqids.Sqids
	}
)

func New(minLength uint8) (*uniqueIDService, error) {

	if minLength == 0 {

		minLength = 10
	}

	engine, err := sqids.New(sqids.Options{MinLength: uint8(minLength)})

	if err != nil {

		return nil, libError.NewInternal(err)
	}

	return &uniqueIDService{
		engine: engine,
	}, nil
}

func (this *uniqueIDService) Serve(input []byte) (string, error) {

	input_length := len(input)
	odd := input_length % ENGINE_UNIT_SIZE
	buffer_capacity := input_length / ENGINE_UNIT_SIZE

	if odd > 0 {

		buffer_capacity++
	}

	spectrum := make([]byte, buffer_capacity*ENGINE_UNIT_SIZE)
	copy(spectrum, input)

	buffer := make([]uint64, buffer_capacity)

	for i := 0; i < buffer_capacity; i++ {

		offset := i * ENGINE_UNIT_SIZE

		buffer[i] = binary.BigEndian.Uint64(spectrum[offset : offset+ENGINE_UNIT_SIZE])
	}

	ret, err := this.engine.Encode(buffer)

	if err != nil {

		return "", libError.NewInternal(err)
	}

	return ret, nil
}

func (this *uniqueIDService) Decode(input string) ([]byte, error) {

	extracted := this.engine.Decode(input)

	ret := make([]byte, len(extracted)*ENGINE_UNIT_SIZE)

	for i, v := range extracted {

		offset := i * ENGINE_UNIT_SIZE

		binary.BigEndian.PutUint64(ret[offset:offset+ENGINE_UNIT_SIZE], v)
	}

	return ret, nil
}
