package variable

import (
	"fmt"
	"testing"
)

func TestNewConstant(t *testing.T) {

	const_str := NewConstain[string]("Hello World!")

	t.Log(const_str)

	fmt.Println(ValueOfConst(const_str))

	const_numeric := NewConstain[int](123)

	t.Log(const_numeric)
}
