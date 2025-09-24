package trie

type (
	base_trie_t[Chain_Node_T comparable, Payload_T any, Chain_T Chainable[Chain_Node_T]] struct {
		base CargoTrie[Chain_Node_T, Payload_T, Chain_T]
	}
)

func (this *base_trie_t[Chain_Node_T, Payload_T, Chain_T]) SetMatcher(matcher IChainNodeValueMatcher[Chain_Node_T]) {
	this.base.SetMatcher(matcher)
}
