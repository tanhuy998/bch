package libAuth

import commonAuth "app/infrastructure/http/common/auth/internal"

func CheckAuthExCludedPath(path string) bool {

	return commonAuth.HasExcluded(path)
}

func CheckAuthAnonymouse(path string) bool {

	return commonAuth.HasAnonymous(path)
}
