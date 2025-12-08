package customerr

import "errors"

var ErrNotFound = errors.New("not found")
var ErrConflict = errors.New("conflict")
var ErrUnauthorized = errors.New("unauthorized")
var ErrInternal = errors.New("internal server Error")
var ErrBadRequest = errors.New("bad request")
var ErrForbidden = errors.New("forbidden")
