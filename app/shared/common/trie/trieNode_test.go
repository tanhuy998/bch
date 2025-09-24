package trie

import (
	"regexp"
	"testing"
)

type (
	r regexp.Regexp
)

func (this *r) Match(nodeValue string, depth int) bool {

	return (*regexp.Regexp)(this).MatchString(nodeValue)
}

func (this *r) StereoTypeFor(nodeValue string, depth int) string {

	return (*regexp.Regexp)(this).String()
}

func (this *r) IsStereoType(pattern string) bool {

	return (*regexp.Regexp)(this).MatchString(pattern)
}

func (this *r) TransformStereotype(pattern string) string {

	return (*regexp.Regexp)(this).String()
}

func TestPattern(t *testing.T) {

	matcher := r(*regexp.MustCompile(`^\{.+\}$`))

	trie := NewDelemiterTrie[bool](DefaultStringSplitter("/"))

	trie.base.SetMatcher(&matcher)

	trie.Merge("GET/a/{foo}/c/{bar}")
	trie.Merge("GET/a/b")
	trie.Merge("POST/a/b")

	switch {
	case trie.Match("GET/a/12314"):
		t.Error("failed 1")
	case trie.Match("GET/a/nsadb"):
		t.Error("failed 2")
	case !trie.Match("GET/a/b"):
		t.Error("failed 3")
	case !trie.Match("GET/a/asd/c/kinisd"):
		t.Error("failed 4")
	}
}

func TestMerge(t *testing.T) {

	root := NewDelemiterTrie[bool](DefaultStringSplitter("/"))

	root.Merge("GET/a/b")
	root.Merge("POST/a/b")

	switch {
	case !root.Match("GET/a/b"):
		t.Error("failed 1")
	case root.Match("POST/a"):
		t.Error("failed 2")
	case root.Match("/a"):
		t.Errorf("failed 3")
	}

}
