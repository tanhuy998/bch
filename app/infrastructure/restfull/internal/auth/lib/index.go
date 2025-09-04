package libAuth

import commonAuth "app/infrastructure/restfull/internal/auth/internal"

func CheckAuthExCludedPath(path string) bool {

	return commonAuth.HasExcluded(path)
}

func CheckAuthAnonymouse(path string) bool {

	return commonAuth.HasAnonymous(path)
}

func CheckNoAuthorize(path string) bool {

	return commonAuth.IsNoAuthorize(path)
}
