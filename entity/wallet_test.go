package entity_test

import (
	"testing"
	"time"

	"github.com/matthewagius/wallet-api/entity"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestInitialiseWallet(t *testing.T) {
	testData := entity.Wallet{
		Id:              1,
		UserId:          123,
		CurrencyCode:    "USD",
		TransactionType: "Credit",
		Amount:          decimal.NewFromFloat(100.00),
		CreatedAt:       time.Now(),
	}
	result, err := entity.Initialise(testData.Id, testData.UserId, testData.CurrencyCode, testData.Amount, testData.TransactionType, testData.CreatedAt)

	assert.Nil(t, err)
	assert.Equal(t, result.Id, testData.Id)
	assert.Equal(t, result.UserId, testData.UserId)
	assert.Equal(t, result.CurrencyCode, testData.CurrencyCode)
	assert.Equal(t, result.Amount, testData.Amount)
	assert.Equal(t, result.TransactionType, testData.TransactionType)
	assert.Equal(t, result.CreatedAt, testData.CreatedAt)
}

func TestInitialiseWalletNevagtiveAmount(t *testing.T) {
	testData := entity.Wallet{
		Id:              1,
		UserId:          123,
		CurrencyCode:    "USD",
		TransactionType: "Credit",
		Amount:          decimal.NewFromFloat(-100.00),
		CreatedAt:       time.Now(),
	}
	result, err := entity.Initialise(testData.Id, testData.UserId, testData.CurrencyCode, testData.Amount, testData.TransactionType, testData.CreatedAt)

	assert.Nil(t, result)
	assert.Equal(t, err, entity.ErrNegativeAmount)
}

func TestInitialiseWalletInavlidWalletId(t *testing.T) {
	testData := entity.Wallet{
		Id:              0,
		UserId:          123,
		CurrencyCode:    "USD",
		TransactionType: "Credit",
		Amount:          decimal.NewFromFloat(100.00),
		CreatedAt:       time.Now(),
	}
	result, err := entity.Initialise(testData.Id, testData.UserId, testData.CurrencyCode, testData.Amount, testData.TransactionType, testData.CreatedAt)

	assert.Nil(t, result)
	assert.Equal(t, err, entity.ErrInvalidEntity)
}

func TestDebitTransactionPostiveBalance(t *testing.T) {
	testData := &entity.Wallet{
		Id:              0,
		UserId:          123,
		CurrencyCode:    "USD",
		TransactionType: "Debit",
		Amount:          decimal.NewFromFloat(100.00),
		CreatedAt:       time.Now(),
	}
	err := entity.ValidateDebitTransaction(decimal.NewFromFloat(10), testData)

	assert.Nil(t, err)
	assert.Equal(t, testData.Amount, decimal.NewFromFloat(90))

}

func TestDebitTransactionNegativeBalance(t *testing.T) {
	testData := &entity.Wallet{
		Id:              0,
		UserId:          123,
		CurrencyCode:    "USD",
		TransactionType: "Debit",
		Amount:          decimal.NewFromFloat(100.00),
		CreatedAt:       time.Now(),
	}
	err := entity.ValidateDebitTransaction(decimal.NewFromFloat(200), testData)

	assert.Equal(t, err, entity.ErrNegativeAmount)
}
