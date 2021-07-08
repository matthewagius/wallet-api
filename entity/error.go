package entity

import "errors"

var ErrRecordNotFound = errors.New("no wallet found with the given ID")
var ErrInvalidEntity = errors.New("invalid entity")
var ErrNegativeAmount = errors.New("wallet amount cannot be negative")
var ErrDebitWallet = errors.New("wallet amount cannot be 0 or of -ve value")
var ErrInvalidTransactionType = errors.New("invalid transaction type, transactions must be either debit or credit")
