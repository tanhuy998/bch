package commonAuthTrie

import (
	"fmt"
)

const (
	UN_STAGED = iota
	EXCLUDE
	ANONYMOUS
)

type (
	trie_node_t struct {
		cur         map[string]*trie_node_t
		kind        int
		depth       uint
		is_param    bool
		is_endpoint bool
	}
)

func (this *trie_node_t) IsEndpoint() bool {

	return this.is_endpoint
}

func (this *trie_node_t) Match(path string) bool {

	_, err := ret_match_enpoint(path, this)

	return err == nil
}

func (this *trie_node_t) Merge(path string) {

	merge(path, this)
}

func (this *trie_node_t) MergeExclude(path string) {

	node := merge(path, this)

	if node.kind != UN_STAGED && node.kind != EXCLUDE {

		panic(
			fmt.Sprintf(
				`path %s was marked as auth anonymous path, could not mark again as auth excluded path`,
				path,
			),
		)
	}

	node.kind = EXCLUDE
}

func (this *trie_node_t) MergeAnonymous(path string) {

	node := merge(path, this)

	if node.kind != UN_STAGED && node.kind != ANONYMOUS {

		panic(
			fmt.Sprintf(
				`path %s was marked as auth excluded path, could not mark again as auth anonymous path`,
				path,
			),
		)
	}

	node.kind = EXCLUDE
}

func (this *trie_node_t) MatchExclude(path string) bool {

	node, err := ret_match_enpoint(path, this)

	if err != nil {

		return false
	}

	return node.kind == EXCLUDE
}

func (this *trie_node_t) MatchAnonymous(path string) bool {

	node, err := ret_match_enpoint(path, this)

	if err != nil {

		return false
	}

	return node.kind == ANONYMOUS
}
