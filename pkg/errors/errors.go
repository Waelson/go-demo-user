package apierrors

type ApiError interface {
	Code() int
	Error() string
}

type apiError struct {
	httpCode int
	message  string
}

func (a *apiError) Code() int {
	return a.httpCode
}

func (a *apiError) Error() string {
	return a.message
}

func New(httpCode int, message string) ApiError {
	return &apiError{
		httpCode: httpCode,
		message:  message,
	}
}

func NewServiceError(message string) ApiError {
	return &apiError{
		httpCode: 400,
		message:  message,
	}
}

func NewInternalError(message string) ApiError {
	return &apiError{
		httpCode: 500,
		message:  message,
	}
}

func NewDatabaseError(message string) ApiError {
	return &apiError{
		httpCode: 500,
		message:  message,
	}
}
