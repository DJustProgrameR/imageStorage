// Package model - domain models
package model

type Err struct {
	Message string `json:"message"`
}

func (e Err) Error() string {
	return e.Message
}

var (
	ErrInvalidAscii = Err{"ascii image is too big"}
	ErrInvalidDescr = Err{"descr is too long"}
	ErrInternal     = Err{"internal server error"}
	ErrNoData       = Err{"no pet uploaded"}
)
