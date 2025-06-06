package filter

func NewFilterGenerator() *FilterGenerator {

	return &FilterGenerator{}
}

func NewConditionFilterGenerator() *ConditionFilterGenerator {

	return new(ConditionFilterGenerator)
}
