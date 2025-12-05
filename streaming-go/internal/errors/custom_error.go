package errors

import "fmt"

type ErrNotFound struct {
	Entity string
	ID     string
}

func (e ErrNotFound) Error() string {
	return fmt.Sprintf("%s with id %s not found", e.Entity, e.ID)
}

type ErrAlreadyExists struct {
	Entity string
	ID     string
}

func (e ErrAlreadyExists) Error() string {
	return fmt.Sprintf("%s with id %s already exists", e.Entity, e.ID)
}
