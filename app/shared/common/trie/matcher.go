package trie

type (
	IChainNodeValueMatcher[Node_Chain_T comparable] interface {
		Match(nodeValue Node_Chain_T, depth int) bool
		IChainNodeSterotypeMatcher[Node_Chain_T]
	}

	IChainNodeSterotypeMatcher[Node_Chain_T comparable] interface {
		IsStereoType(pattern Node_Chain_T) bool
		TransformStereotype(pattern Node_Chain_T) Node_Chain_T
		StereoTypeFor(nodeValue Node_Chain_T, depth int) Node_Chain_T
	}
)
