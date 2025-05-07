package filter

func NewFilterGenerator() *filter_generator {

	return &filter_generator{}
}

func NewConditionFilterGenerator() *condition_filter_generator {

	return new(condition_filter_generator)
}
