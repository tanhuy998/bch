package trie

type (
	trie_node_t[Chain_Node_T comparable, Payload_T any, Chain_T Chainable[Chain_Node_T]] struct {
		m          map[Chain_Node_T]*trie_node_t[Chain_Node_T, Payload_T, Chain_T]
		stereotype map[Chain_Node_T]Chain_Node_T
		payload    *Payload_T
		depth      int
	}
)

func newTrieNode[Chain_Node_T comparable, Payload_T any, Chain_T Chainable[Chain_Node_T]](
	depth int,
) *trie_node_t[Chain_Node_T, Payload_T, Chain_T] {

	return &trie_node_t[Chain_Node_T, Payload_T, Chain_T]{
		depth: depth,
		m:     make(map[Chain_Node_T]*trie_node_t[Chain_Node_T, Payload_T, Chain_T]),
	}
}

func (this *trie_node_t[Chain_Node_T, Payload_T, Chain_T]) _map(strereotype Chain_Node_T, trieLookedupValue Chain_Node_T) {

	if this.stereotype == nil {

		this.stereotype = make(map[Chain_Node_T]Chain_Node_T)
	}

	this.stereotype[strereotype] = trieLookedupValue
}

func (this *trie_node_t[Chain_Node_T, Payload_T, Chain_T]) Find(
	chain Chain_T, matcher IChainNodeValueMatcher[Chain_Node_T],
) (payload Payload_T, match bool) {

	switch node, _match := this.find(chain, matcher); {
	case _match:
		return *node.payload, true
	default:
		return
	}
}

func (this *trie_node_t[Chain_Node_T, Payload_T, Chain_T]) find(
	chain Chain_T, matcher IChainNodeValueMatcher[Chain_Node_T],
) (node *trie_node_t[Chain_Node_T, Payload_T, Chain_T], match bool) {

	switch _node, remain := _lookup(this, chain, matcher); {
	case len(remain) == 0 && _node != nil && _node.payload != nil:
		return _node, true
	default:
		return
	}
}

func (this *trie_node_t[Chain_Node_T, Payload_T, Chain_T]) merge(
	chain Chain_T, matcher IChainNodeValueMatcher[Chain_Node_T],
) (node *trie_node_t[Chain_Node_T, Payload_T, Chain_T]) {

	lastMatchNode, remainingChain := _lookup_append(this, chain, matcher)

	switch len(remainingChain) {
	case 0:
		node = lastMatchNode
	default:
		node = __append(lastMatchNode, remainingChain, matcher)
	}

	return
}

func (this *trie_node_t[Chain_Node_T, Payload_T, Chain_T]) Set(
	chain Chain_T, payload Payload_T, matcher IChainNodeValueMatcher[Chain_Node_T],
) {

	node := this.merge(chain, matcher)

	node.payload = &payload
}

func (this *trie_node_t[Chain_Node_T, Payload_T, Chain_T]) Merge(chain Chain_T, matcher IChainNodeValueMatcher[Chain_Node_T]) {

	node := this.merge(chain, matcher)

	node.payload = new(Payload_T)
}

// Deep dive into trie hierachy until the input chain remains nothing.
// This fucntion returns the last match node. If remain is not empty, last match node
// belongs to the nearest merged to the input, remain is the unmerged discriminations.
func _lookup[Chain_Node_T comparable, Payload_T any, Chain_T Chainable[Chain_Node_T]](
	entryNode *trie_node_t[Chain_Node_T, Payload_T, Chain_T], chain Chain_T, matcher IChainNodeValueMatcher[Chain_Node_T],
) (lastMatchNode *trie_node_t[Chain_Node_T, Payload_T, Chain_T], remain Chain_T) {

	const (
		FIRST = iota
		SECOND
	)

	var (
		//curr_node *trie_node_t[Chain_Node_T, Payload_T, Chain_T]
		f_matcher IChainNodeValueMatcher[Chain_Node_T] = matcher
	)

	remain = chain
	lastMatchNode = entryNode

	for range len(remain) {

		var (
			lookedupValue Chain_Node_T = remain[FIRST]
		)

	__LABEL_MAP_LOOKUP__:
		switch nextNode, exists := lastMatchNode.m[lookedupValue]; {
		case exists:
			lastMatchNode = nextNode
			remain = remain[SECOND:]
			continue
		case f_matcher != nil && !f_matcher.Match(lookedupValue, lastMatchNode.depth):
			ste := f_matcher.StereoTypeFor(lookedupValue, lastMatchNode.depth)
			lookedupValue = lastMatchNode.stereotype[ste]
			f_matcher = nil
			goto __LABEL_MAP_LOOKUP__
		default:
			return
		}
	}

	return //curr_node, remain
}

func _lookup_append[Chain_Node_T comparable, Payload_T any, Chain_T Chainable[Chain_Node_T]](
	entryNode *trie_node_t[Chain_Node_T, Payload_T, Chain_T], chain Chain_T, matcher IChainNodeSterotypeMatcher[Chain_Node_T],
) (lastMatchNode *trie_node_t[Chain_Node_T, Payload_T, Chain_T], remain Chain_T) {

	const (
		FIRST = iota
		SECOND
	)

	var (
		//curr_node *trie_node_t[Chain_Node_T, Payload_T, Chain_T]
		f_matcher IChainNodeSterotypeMatcher[Chain_Node_T] = matcher
	)

	remain = chain
	lastMatchNode = entryNode

	for range len(remain) {

		var (
			lookedupValue Chain_Node_T = remain[FIRST]
		)

	__LABEL_MAP_LOOKUP__:
		switch nextNode, exists := lastMatchNode.m[lookedupValue]; {
		case exists:
			lastMatchNode = nextNode
			remain = remain[SECOND:]
			continue
		case f_matcher != nil && f_matcher.IsStereoType(lookedupValue):
			ste := matcher.TransformStereotype(lookedupValue)
			lookedupValue = lastMatchNode.stereotype[ste]
			f_matcher = nil
			goto __LABEL_MAP_LOOKUP__
		default:
			return
		}
	}

	return
}

func __append[Chain_T Chainable[Chain_Node_T], Chain_Node_T comparable, Payload_T any](
	entryNode *trie_node_t[Chain_Node_T, Payload_T, Chain_T], patternChain Chain_T, matcher IChainNodeSterotypeMatcher[Chain_Node_T],
) (lastNode *trie_node_t[Chain_Node_T, Payload_T, Chain_T]) {

	const (
		FIRST = iota
		SECOND
	)

	var curr_node *trie_node_t[Chain_Node_T, Payload_T, Chain_T]

	curr_node = entryNode
	remain := patternChain

	for range len(remain) {

		nextNode := newTrieNode[Chain_Node_T, Payload_T, Chain_T](curr_node.depth + 1)

		var storedValue Chain_Node_T = remain[FIRST]

		switch {
		case matcher != nil && matcher.IsStereoType(remain[FIRST]):
			ste := matcher.TransformStereotype(remain[FIRST])
			curr_node._map(ste, remain[FIRST])
		}

		curr_node.m[storedValue] = nextNode
		curr_node = nextNode
		remain = remain[SECOND:]
	}

	return curr_node
}
