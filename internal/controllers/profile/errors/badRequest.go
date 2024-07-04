package errors

import (
	"fmt"
	"net/http"
)

type Error struct {
	status int
	msg    string
	args   []any
}

func (e Error) Error() string {
	if len(e.args) == 0 {
		return e.msg
	}
	return fmt.Sprintf(e.msg, e.args...)
}

func (e Error) Status() int {
	return e.status
}

var (
	BadRequest     = Error{http.StatusBadRequest, "could not parse request body", []any{}}
	InternalServer = Error{http.StatusInternalServerError, "something went wrong", []any{}}
)
