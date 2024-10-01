package common

type (
	HeadResponse struct {
	}
)

func (this *HeadResponse) IsNotContent() bool {

	return true
}
