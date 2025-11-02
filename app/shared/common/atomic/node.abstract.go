package atomic

type (
	IRaceValue[Value_T any] interface {
		setRaceValue(v Value_T)
	}
)

func SetRaceValueOf[Value_T any](setter IRaceValue[Value_T], v Value_T) {

	setter.setRaceValue(v)
}

type (
	abstract_node_t[Value_T any] struct {
		node_mutex_t
		val Value_T
	}
)

func (this *abstract_node_t[Value_T]) isNode() {}

func (this *abstract_node_t[Value_T]) Value() Value_T {

	this.RLock()
	defer this.RUnlock()

	return this.val
}

func (this *abstract_node_t[Value_T]) SetValue(v Value_T) {

	this.Lock()
	defer this.Unlock()

	this.val = v
}

func (this *abstract_node_t[Value_T]) setRaceValue(v Value_T) {

	this.val = v
}
