package libCommon

import (
	"reflect"
	"testing"
)

func Test_EvaluateIndirectNilValueByPointer(t *testing.T) {

	type T struct {
		prop int
	}

	var (
		// indirectPtr **T
		ptr *T
	)

	//indirectPtr = &ptr

	EvaluateIndirectValueOfPointer[*T](&ptr)

	if ptr == nil {

		t.Fatal("failed")
		return
	}

	t.Log("ok")
}

func Test_EvaluateIndirectNilValueByReflection(t *testing.T) {

	type T struct {
		prop int
	}

	var (
		ptr *T
	)

	EvaluateIndirectValueBy(
		reflect.Indirect(
			reflect.ValueOf(&t),
		),
	)

	if ptr == nil {

		t.Fatal("failed")
		return
	}

	t.Log("ok")
}
