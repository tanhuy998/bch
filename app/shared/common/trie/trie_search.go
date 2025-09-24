package trie

type (
	SearchTrie[Chain_Node_T comparable, Chain_T Chainable[Chain_Node_T]] struct {
		base_trie_t[Chain_Node_T, bool, Chain_T]
	}
)

func NewSearchTrie[Chain_Node_T comparable, Chain_T Chainable[Chain_Node_T]]() *SearchTrie[Chain_Node_T, Chain_T] {

	ret := new(SearchTrie[Chain_Node_T, Chain_T])
	ret.init()
	return ret
}

func (this *SearchTrie[Chain_Node_T, Chain_T]) init() {

	this.base.init()
}

func (this *SearchTrie[Chain_Node_T, Chain_T]) Merge(chain Chain_T) {

	this.base.MergeWithPayload(
		chain, true,
	)
}

func (this *SearchTrie[Chain_Node_T, Chain_T]) Find(chain Chain_T) bool {

	_, ret := this.base.Find(chain)

	return ret
}
