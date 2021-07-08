package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type Wallet struct {
	Id           uint            `json:"walletid"`
	UserId       uint            `json:"userId"`
	CurrencyCode string          `json:"currencyCode"`
	Amount       decimal.Decimal `json:"amount"`
	CreatedAt    time.Time       `json:"createdat"`
	UpdatedAt    time.Time       `json:"updatedat"`
}

type TransactionDetail struct {
	UserId       uint            `json:"userId"`
	CurrencyCode string          `json:"currencyCode"`
	Amount       decimal.Decimal `json:"amount"`
}
