// Package model - domain models
package model

import "strings"

// Err domain model
type Err struct {
	Message string `json:"message"`
}

// Error provides error text
func (e Err) Error() string {
	return e.Message
}

// Is checks if err contains model.Err content
func (e Err) Is(err error) bool {
	return strings.Contains(err.Error(), e.Message)
}

// Default errs
var (
	ErrInvalidASCII = Err{"ascii image is too big"}
	ErrInvalidDescr = Err{"descr is too long"}
	ErrInternal     = Err{"internal server error"}
	ErrNoData       = Err{"no pet uploaded"}
)
