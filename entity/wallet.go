package entity

import (
	"time"

	"github.com/shopspring/decimal"
)

type Wallet struct {
	Id              uint `gorm:"primaryKey, AUTO_INCREMENT"`
	UserId          uint
	CurrencyCode    string
	Amount          decimal.Decimal
	CreatedAt       time.Time
	UpdatedAt       time.Time
	TransactionType string `gorm:"-"`
}

func UpdateWallet(walletId uint, userId uint, currencyCode string, amount decimal.Decimal, transactionType string, createdAt time.Time) (*Wallet, error) {
	wallet := &Wallet{
		Id:              walletId,
		UserId:          userId,
		CurrencyCode:    currencyCode,
		TransactionType: transactionType,
		Amount:          amount,
		CreatedAt:       createdAt,
	}

	err := wallet.Validate()
	if err != nil {
		return nil, err
	}

	return wallet, nil
}

func (w *Wallet) Validate() error {

	if w.Amount.IsNegative() {
		return ErrNegativeAmount
	}

	if w.Id == 0 {
		return ErrInvalidEntity
	}
	return nil
}
