package servicePort

/*
	Adapters solve the problem of circular dependencies when
	cross domain/package/namespace services calling each other.

	Apdapters must be register before all services.
*/

type (
	IAdapter[T any] interface {
		Adaptee() *T
	}
)
