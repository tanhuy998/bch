package libError

func IsAcceptable(target error, exceptions ...error) bool {

	if target == nil {

		return true
	}

	for _, exceptErr := range exceptions {

		if exceptErr == nil {

			continue
		}

		if target == exceptErr {

			return true
		}

	}

	return false
}
