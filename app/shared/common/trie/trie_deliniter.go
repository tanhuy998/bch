package trie

type (
	delimiter_trie_t[Payload_T any] struct {
		base_trie_t[string, Payload_T, []string]
		splitter IStringSplitter
	}
)

func NewDelemiterTrie[Payload_T any](splitter IStringSplitter) *delimiter_trie_t[Payload_T] {

	if splitter == nil {

		panic("delimiter splitter must not be nil")
	}

	ret := new(delimiter_trie_t[Payload_T])
	ret.splitter = splitter
	ret.base.init()

	return ret
}

func (this *delimiter_trie_t[Payload_T]) Merge(str string) {

	exploded := this.splitter.Split(str)
	this.base.Merge(exploded)
}

func (this *delimiter_trie_t[Payload_T]) Find(str string) (payload Payload_T, match bool) {

	exploded := this.splitter.Split(str)

	return this.base.Find(exploded)
}

func (this *delimiter_trie_t[Payload_T]) Match(str string) bool {

	exploded := this.splitter.Split(str)

	return this.base.Match(exploded)
}
