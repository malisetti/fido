package storage

import "errors"

type SystemError struct {
}

func (err *SystemError) Error() error {
	return errors.New("System error")
}
