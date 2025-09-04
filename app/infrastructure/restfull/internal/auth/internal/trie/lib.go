package commonAuthTrie

import (
	libCommon "app/internal/lib/common"
	"fmt"
	"regexp"
	"strings"
)

var (
	pathParam      = regexp.MustCompile(`\{.+\}`)
	errNoMatchNode = fmt.Errorf("no match trie node")
)

func New() trie_node_t {

	return newTrieNode(nil)
}

func newTrieNode(ref *trie_node_t) (ret trie_node_t) {

	ret = trie_node_t{}
	ret.cur = make(map[string]*trie_node_t)

	if ref == nil {

		return
	}

	ret.depth = ref.depth + 1

	return ret
}

func merge(path string, ref *trie_node_t) *trie_node_t {

	pathParts := strings.Split(path, "/")
	currentNode := ref

	for _, part := range pathParts {

		if part == "" {

			continue
		}

		if pathParam.Match([]byte(part)) {

			currentNode.is_param = true
			part = "{}"
		}

		if currentNode.cur == nil {

			currentNode.cur = make(map[string]*trie_node_t)
		}

		if _, has := currentNode.cur[part]; !has {

			currentNode.cur[part] = libCommon.PointerPrimitive(newTrieNode(currentNode))
		}

		currentNode = currentNode.cur[part]
	}

	currentNode.is_endpoint = true
	return currentNode
}

func ret_match(path string, ref *trie_node_t) (ret trie_node_t, err error) {

	pathParts := strings.Split(path, "/")

	currentNode := ref

	for _, part := range pathParts {

		if len(currentNode.cur) == 0 {

			err = errNoMatchNode
			return
		}

		node, has := currentNode.cur[part]

		if has {

			currentNode = node
			continue
		}

		if !currentNode.is_param {

			err = errNoMatchNode
			return

		} else {

			currentNode = currentNode.cur["{}"]
		}
	}

	return *currentNode, nil
}

func ret_match_enpoint(path string, ref *trie_node_t) (ret trie_node_t, err error) {

	node, err := ret_match(path, ref)

	if err != nil {

		return
	}

	if !node.is_endpoint {

		err = errNoMatchNode
		return
	}

	return node, nil
}
