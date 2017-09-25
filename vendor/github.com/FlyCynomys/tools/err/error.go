package err

import (
	"encoding/json"
	"errors"
	"fmt"
)

type ErrorCode struct {
	Code int
	error
}

func (e *ErrorCode) Error() string {
	return fmt.Sprintf("code : %d ,error : %s ", e.Code, e.error.Error())
}

func (e *ErrorCode) Json() []byte {
	data, _ := json.Marshal(e)
	return data
}

func New(code int, description string) *ErrorCode {
	return &ErrorCode{
		Code:  code,
		error: errors.New(description),
	}
}
