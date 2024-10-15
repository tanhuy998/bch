package requestInput

type (
	RangePaginateInput struct {
		PageNumber uint64 `url:"page"`
		PageSize   uint64 `url:"size"`
	}
)

func (this *RangePaginateInput) GetPageNumber() uint64 {

	return this.PageNumber
}

func (this *RangePaginateInput) GetPageSize() uint64 {

	return this.PageSize
}
