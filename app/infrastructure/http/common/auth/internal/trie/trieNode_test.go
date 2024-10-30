package commonAuthTrie_test

import (
	commonAuthTrie "app/infrastructure/http/common/auth/internal/trie"
	"testing"
)

func TestMain(t *testing.T) {

	root := commonAuthTrie.New()

	root.Merge("GET/a/b")
	root.Merge("POST/a/b")

	if !root.Match("GET/a/b") {

		t.Error("failed 1")
	}

	if root.Match("POST/a") {

		t.Error("failed 2")
	}

	if root.Match("/a") {

		t.Errorf("failed 3")
	}
}

func TestParam(t *testing.T) {

	root := commonAuthTrie.New()

	root.Merge("GET/a/b")
	root.Merge("POST/a/b")
	root.Merge("POST/a/{p}")

	if !root.Match("GET/a/b") {

		t.Error("failed 1")
	}

	if root.Match("POST/a") {

		t.Error("failed 2")
	}

	if root.Match("/a") {

		t.Errorf("failed 3")
	}

	if !root.Match("POST/a/123") {

		t.Error("failed 4")
	}

	if root.Match("GET/a/12") {

		t.Error("failed 5")
	}

	if !root.Match("POST/a/b") {

		t.Error("failed 6")
	}
}
