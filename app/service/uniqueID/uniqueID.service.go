package uniqueIDService

import (
	uniqueIDServicePort "app/adapter/uniqueID"
	"encoding/binary"

	"github.com/sqids/sqids-go"
)

const (
	ENGINE_UNIT_SIZE = 8
)

type (
	IUniqueIDGenerator = uniqueIDServicePort.IUniqueIDGenerator

	uniqueIDService struct {
		engine *sqids.Sqids
	}
)

func New(minLength uint) (*uniqueIDService, error) {

	engine, err := sqids.New(sqids.Options{MinLength: 10})

	if err != nil {

		return nil, err
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

	buffer := []uint64{}

	for i := 0; i < input_length; i += ENGINE_UNIT_SIZE {

		var (
			i_end int
		)

		if input_length-i == odd {

			i_end = i + odd

		} else {

			i_end = i + ENGINE_UNIT_SIZE
		}

		buffer = append(buffer, binary.BigEndian.Uint64(input[i:i_end]))
	}

	return this.engine.Encode(buffer)
}
