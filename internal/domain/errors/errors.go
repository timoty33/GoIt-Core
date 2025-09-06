package errors

import (
	"fmt"
)

type GoItError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Err     error  `json:"-"`
}

func (e *GoItError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("[%s] %s: %v", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

func New(code, message string, err error) GoItError {
	return GoItError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

