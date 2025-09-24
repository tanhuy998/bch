package env

import (
	"app/port/envConsumerServicePort"
)

type (
	ENVConsumer struct {
	}
)

func (this *ENVConsumer) Get(key string) envConsumerServicePort.ITypeAssertionValue {

	return env_assertion_val_t("")
}
