package errorMap

import libError "app/internal/lib/error"

func WrapAndExportAsInternalError(errs ...error) error {

	return libError.NewInternal(errs...)
}
