package inner

import (
	"app/internal/db/query"
	"fmt"
	"strings"
)

type (
	self_reference_t string
)

func (this self_reference_t) GetQuerySelfReference() string {

	return this.String()
}

func (this self_reference_t) Origin() string {

	return string(this)
}

func (this self_reference_t) String() string {

	return fmt.Sprintf(`$%s`, string(this))
}

func AssertSelfReference(in string) (ret query.IQuerySelfReference, ok bool) {

	switch s, ok := strings.CutPrefix(in, "$"); {
	case ok:
		return self_reference_t(s), ok
	default:
		return self_reference_t(""), false
	}

}

func Self(field string) query.IQuerySelfReference {

	return self_reference_t(field)
}
