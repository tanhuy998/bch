package errorMap

type (
	ICodeError interface {
		error
		GetCode() ERROR_CODE_T
	}

	code_error_t struct {
		code ERROR_CODE_T
	}
)

func (this code_error_t) Error() string {

	return ""
}

func (this code_error_t) GetCode() ERROR_CODE_T {

	return this.code
}
