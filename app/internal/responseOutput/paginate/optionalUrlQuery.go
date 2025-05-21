package paginateOutput

type (
	OptionalUrlQuery map[string]interface{}
)

func (this *OptionalUrlQuery) init() {

	if *this != nil {
		return
	}

	*this = make(OptionalUrlQuery)
}

func (this *OptionalUrlQuery) Set(
	key string, val interface{},
) {

	this.init()

	(*this)[key] = val
}
