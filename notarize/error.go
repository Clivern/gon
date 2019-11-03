package notarize

import (
	"fmt"

	"github.com/hashicorp/go-multierror"
)

// Error is the error structure generated by the notarization tool.
type Error struct {
	Code     int64             `plist:"code"`
	Message  string            `plist:"message"`
	UserInfo map[string]string `plist:"userInfo"`
}

// Errors is a list of error and also implements error.
type Errors []Error

// Error implements error
func (err Error) Error() string {
	return fmt.Sprintf("%s (%d)", err.Message, err.Code)
}

// Error implements error
func (err Errors) Error() string {
	if len(err) == 0 {
		return "no errors"
	}

	var result error
	for _, e := range err {
		result = multierror.Append(result, e)
	}

	return result.Error()
}

// ContainsCode returns true if the errors list has an error with the given code.
func (err Errors) ContainsCode(code int64) bool {
	for _, e := range err {
		if e.Code == code {
			return true
		}
	}

	return false
}
