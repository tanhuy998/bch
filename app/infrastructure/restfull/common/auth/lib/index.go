package libAuth

import commonAuth "app/infrastructure/restfull/common/auth/internal"

func CheckAuthExCludedPath(path string) bool {

	return commonAuth.HasExcluded(path)
}

func CheckAuthAnonymouse(path string) bool {

	return commonAuth.HasAnonymous(path)
}
