package storage

import "errors"

const SerialVersionUID = 1L

type SystemError struct {
}

func (err *SystemError) Error() error  {
    return errors.New("System error")    
}