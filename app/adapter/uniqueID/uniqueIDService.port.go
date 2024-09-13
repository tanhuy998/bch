package uniqueIDServicePort

type (
	IUniqueIDGenerator interface {
		Serve(bytes []byte) (string, error)
	}
)
