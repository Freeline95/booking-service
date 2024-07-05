package error

type BusinessError struct {
	code        string
	description string
}

func NewBusinessError(code string, description string) BusinessError {
	return BusinessError{
		code:        code,
		description: description,
	}
}

func (be BusinessError) Error() string {
	return be.description
}

func (be BusinessError) Code() string {
	return be.code
}
