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

func Initialise(walletId uint, userId uint, currencyCode string, amount decimal.Decimal, transactionType string, createdAt time.Time) (*Wallet, error) {
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

//check that amount is positive and a positive ID greater than 0 has been passed
func (w *Wallet) Validate() error {
	if w.Amount.IsNegative() {
		return ErrNegativeAmount
	}

	if w.Id <= 0 {
		return ErrInvalidEntity
	}
	return nil
}

func ValidateDebitTransaction(newAmount decimal.Decimal, data *Wallet) error {
	var result = data.Amount.Sub(newAmount)
	if result.IsNegative() {
		return ErrNegativeAmount
	}
	data.Amount = result
	return nil
}
