package storage

import "errors"

const SerialVersionUID = 1L

type DuplicateKeyError struct {
}

func (err *DuplicateKeyError) Error() error {
    return errors.New("Duplicate key error")
}
