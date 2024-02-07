package wraperrors

import (
	"errors"
	"fmt"
)

type ErrWrap struct {
	errs []error
}

func NewErrWrap(errs ...error) error {
	return &ErrWrap{errs}
}

func (ew *ErrWrap) Error() string {
	var errorsText []interface{}

	for _, err := range ew.errs {
		errorsText = append(errorsText, err.Error())
	}

	return fmt.Sprintln(errorsText...)
}

func (ew *ErrWrap) Unwrap() error {
	return ew.errs[0]
}

func (ew *ErrWrap) Is(target error) bool {
	for _, err := range ew.errs {
		if errors.Is(err, target) {
			return true
		}
	}

	return false
}
