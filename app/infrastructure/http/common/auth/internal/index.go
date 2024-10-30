package commonAuth

import (
	commonAuthTrie "app/infrastructure/http/common/auth/internal/trie"
	libCommon "app/internal/lib/common"
)

var (
	auth_exlcuded = libCommon.PointerPrimitive(commonAuthTrie.New())
)

func ExcludePath(path string) {

	auth_exlcuded.MergeExclude(path)
}

func HasExcluded(path string) bool {

	return auth_exlcuded.MatchExclude(path)
}

func MarkAnonymous(path string) {

	auth_exlcuded.MergeAnonymous(path)
}

func HasAnonymous(path string) bool {

	return auth_exlcuded.MatchAnonymous(path)
}
