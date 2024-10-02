package responseOutput

/*
this package define interfaces that helps controllers determine which status
the app response to the client.
*/

type (
	INoContentOutput interface {
		IsNotContent() bool
	}

	NoContent struct {
	}
)

func (this *NoContent) IsNotContent() bool {

	return true
}

type (
	ICreatedOutput interface {
		IsCreatedStatus() bool
	}

	CreatedResponse struct {
	}
)

func (this CreatedResponse) IsCreatedStatus() bool {

	return true
}
