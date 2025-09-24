package inputQueryFilterPort

import "time"

type (
	ISingleRangeDateFilter interface {
		GetFilteredDateField() string
		GetFilteredDateBefore() *time.Time
		GetFilteredDateAfter() *time.Time
	}
)
