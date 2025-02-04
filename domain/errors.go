package domain

import "errors"

var (
	ErrInsufficientBalance = errors.New("insufficient balance")
	ErrAccountNotFound     = errors.New("account not found")
)
