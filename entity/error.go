package entity

import "errors"

var ErrRecordNotFound = errors.New("record not found")
var ErrInvalidEntity = errors.New("invalid entity")
var ErrNegativeAmount = errors.New("wallet amount cannot be negative")
var ErrDebitWallet = errors.New("wallet amount cannot be 0 or of -ve value")
