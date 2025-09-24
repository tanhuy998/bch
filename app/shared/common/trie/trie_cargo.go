package trie

import (
	"sync"
)

type (
	Chainable[T comparable] interface {
		~[]T
	}
)

type (
	CargoTrie[Chain_Node_T comparable, Payload_T any, Chain_T Chainable[Chain_Node_T]] struct {
		sync.RWMutex
		matcher IChainNodeValueMatcher[Chain_Node_T]
		root    *trie_node_t[Chain_Node_T, Payload_T, Chain_T]
	}
)

func NewCargoTrie[Chain_Node_T comparable, Payload_T any, Chain_T Chainable[Chain_Node_T]](
	matcher IChainNodeValueMatcher[Chain_Node_T],
) *CargoTrie[Chain_Node_T, Payload_T, Chain_T] {

	return &CargoTrie[Chain_Node_T, Payload_T, Chain_T]{
		root:    newTrieNode[Chain_Node_T, Payload_T, Chain_T](0),
		matcher: matcher,
	}
}

func (this *CargoTrie[Chain_Node_T, Payload_T, Chain_T]) init() {

	if this.root != nil {

		return
	}

	this.root = newTrieNode[Chain_Node_T, Payload_T, Chain_T](0)
}

func (this *CargoTrie[Chain_Node_T, Payload_T, Chain_T]) SetMatcher(matcher IChainNodeValueMatcher[Chain_Node_T]) {

	this.matcher = matcher
}

func (this *CargoTrie[Chain_Node_T, Payload_T, Chain_T]) Merge(chain Chain_T) {

	this.Lock()
	defer this.Unlock()

	this.root.Merge(chain, this.matcher)
}

func (this *CargoTrie[Chain_Node_T, Payload_T, Chain_T]) MergeWithPayload(chain Chain_T, payload Payload_T) {

	this.Lock()
	defer this.Unlock()

	this.root.Set(chain, payload, this.matcher)
}

func (this *CargoTrie[Chain_Node_T, Payload_T, Chain_T]) Match(chain Chain_T) bool {

	this.RLock()
	defer this.RUnlock()

	_, ret := this.Find(chain)

	return ret
}

func (this *CargoTrie[Chain_Node_T, Payload_T, Chain_T]) Find(chain Chain_T) (payload Payload_T, matched bool) {

	this.RLock()
	defer this.RUnlock()

	return this.root.Find(chain, this.matcher)
}

func (this *CargoTrie[Chain_Node_T, Payload_T, Chain_T]) Set(chain Chain_T, payload Payload_T) {

	this.Lock()
	defer this.Unlock()

	this.root.Set(chain, payload, this.matcher)
}
