package lib

type (
	Expression_Map[T any] struct {
		cur string
		m   map[string]T
	}
)

func (this *Expression_Map[T]) GetMap() map[string]T {

	return this.m
}

func (this *Expression_Map[T]) init() {

	if this.m != nil {

		return
	}

	this.m = make(map[string]T)
}

func (this *Expression_Map[T]) Set(key string) *Expression_Map[T] {

	this.cur = key

	return this
}

func (this *Expression_Map[T]) Value(val T) {

	key := this.cur

	if key == "" {

		return
	}

	this.init()

	this.m[key] = val
}
