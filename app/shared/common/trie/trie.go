package trie

import (
	"unicode/utf8"
)

type (
	Trie struct {
		SearchTrie[string, []string]
	}
)

func NewTrie() *Trie {

	ret := new(Trie)
	ret.init()
	return ret
}

func (this *Trie) init() {

	this.base.init()
}

func (this *Trie) Merge(str string) {

	this.base.Merge(
		endChainString(str),
	)
}

func (this *Trie) Find(str string) bool {

	return this.base.Match(endChainString(str))
}

func endChainString(str string) []string {

	chain := make([]string, utf8.RuneCountInString(str))

	for i, r := range str {

		chain[i] = string(r)
	}

	return chain
}
