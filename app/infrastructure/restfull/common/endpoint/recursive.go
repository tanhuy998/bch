package endpoint

type (
	IRecursiveAPICurator interface {
		Child() []IAPICurator
	}
)
