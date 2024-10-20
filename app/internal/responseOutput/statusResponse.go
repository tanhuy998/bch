package responseOutput

type (
	IHTTPStatusResponse interface {
		SetHTTPStatus(code int)
		GetHTTPStatus() int
	}
	HTTPStatusResponse struct {
		status int
	}
)

func (this *HTTPStatusResponse) SetHTTPStatus(code int) {

	this.status = code
}

func (this *HTTPStatusResponse) GetHTTPStatus() int {

	if this.status == 0 {

		return 200
	}

	return this.status
}
