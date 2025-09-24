package filter

func NewFilterGenerator() *filter_t /**FilterGenerator*/ {

	//return &FilterGenerator{}

	return new(filter_t)
}

func NewConditionFilterGenerator() *ConditionFilterGenerator {

	return new(ConditionFilterGenerator)
}
