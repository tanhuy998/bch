package requestInput

type (
	IgnorablePaginator struct {
		IsIgnorePaginate bool `url:"p_ignore"`
	}
)

func (this IgnorablePaginator) IsIgnorePagination() bool {

	return this.IsIgnorePaginate
}
