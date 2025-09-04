package commonAuth

import (
	commonAuthTrie "app/infrastructure/restfull/internal/auth/internal/trie"
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

func MarkNoAuthorize(path string) {

	auth_exlcuded.MergeNoAuthorize(path)
}

func IsNoAuthorize(path string) bool {

	return auth_exlcuded.MatchNoAuthorize(path)
}
