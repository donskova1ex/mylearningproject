package internal

import (
	"errors"
)

func ErrEntityNotFound() error {
	return errors.New("entity not found")
}

func ErrEntityGetByUUID() error {
	return errors.New("can not get entity")
}

func ErrReadRows() error {
	return errors.New("can not read rows")
}
