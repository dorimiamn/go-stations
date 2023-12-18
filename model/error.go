package model

import "fmt"

type ErrNotFound struct {
	Message string
}

func (e ErrNotFound) Error() string {
	return fmt.Sprintf("Not Found: %s", e.Message)
}
