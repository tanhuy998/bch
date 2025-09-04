package annotation

import (
	"app/shared/common/variable"
	"fmt"
	"reflect"
	"time"
)

var signature struct {
	_ string
	variable.NoCopy
} = struct {
	_ string
	variable.NoCopy
}{
	fmt.Sprintf(`%s %s`, time.Now(), "(base annotation)"), variable.NoCopy{},
}

type (
	IAnnotation interface {
		Prove() (evidenc interface{})
	}
	Annotation struct{}
)

func (Annotation) Prove() (evidence interface{}) {

	return signature
}

func AssertAnnotation(target reflect.Value) {

	switch v := target.Interface().(type) {
	case IAnnotation:
		reflect.DeepEqual(v.Prove(), signature)
	default:
		panic("expect the type that embed the Annotation type")
	}
}
