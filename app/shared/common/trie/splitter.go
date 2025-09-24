package trie

import "strings"

type (
	IStringSplitter interface {
		Split(str string) []string
	}
)

type (
	DefaultStringSplitter string
)

func (delimiter DefaultStringSplitter) Split(str string) []string {

	return strings.Split(str, string(delimiter))
}
