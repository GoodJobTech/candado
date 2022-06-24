package errors

import "errors"

var (
	ErrAlreadyLocked    = errors.New("Lock is already acquired")
	ErrAlreadyUnlocked  = errors.New("Lock is not acquired")
	ErrLockDoesntExists = errors.New("Lock does not exists")
)
