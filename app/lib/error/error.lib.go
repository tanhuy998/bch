package libError

func IsErrorExcepts(target error, exceptions ...error) bool {

	if target == nil {

		return false
	}

	for _, errVal := range exceptions {

		if target == errVal {

			return true
		}

	}

	return true
}
