package errors

type HttpError struct {
	S    string `json:"message"`
	Code int    `json:"code"`
}

func (e *HttpError) Error() string {
	return e.S
}

func NewHttp(statusCode int, message string) error {
	return &HttpError{message, statusCode}
}

func GetStatusCode(error interface{}, defaultValue int) int {
	if e, ok := error.(*HttpError); ok {
		return e.Code
	}

	return defaultValue
}