package util

import (
	"time"

	"github.com/matthewagius/wallet-api/entity"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

func InsertSampleData(db *gorm.DB) {
	var count int64
	if db != nil {
		db.Table("wallets").Count(&count)
	}

	if count == 0 {
		data := []entity.Wallet{
			{UserId: 1, CurrencyCode: "EUR", Amount: decimal.NewFromFloat(202.65), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{UserId: 2, CurrencyCode: "GBP", Amount: decimal.NewFromFloat(30.20), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{UserId: 3, CurrencyCode: "USD", Amount: decimal.NewFromFloat(60.00), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{UserId: 4, CurrencyCode: "EUR", Amount: decimal.NewFromFloat(675.00), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{UserId: 5, CurrencyCode: "GBP", Amount: decimal.NewFromFloat(100.00), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{UserId: 6, CurrencyCode: "USD", Amount: decimal.NewFromFloat(23.69), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{UserId: 7, CurrencyCode: "EUR", Amount: decimal.NewFromFloat(46.21), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{UserId: 8, CurrencyCode: "GBP", Amount: decimal.NewFromFloat(82.07), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{UserId: 9, CurrencyCode: "USD", Amount: decimal.NewFromFloat(0.00), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{UserId: 10, CurrencyCode: "EUR", Amount: decimal.NewFromFloat(200.00), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{UserId: 11, CurrencyCode: "GBP", Amount: decimal.NewFromFloat(52.25), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{UserId: 12, CurrencyCode: "USD", Amount: decimal.NewFromFloat(66.67), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{UserId: 13, CurrencyCode: "EUR", Amount: decimal.NewFromFloat(12.67), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{UserId: 14, CurrencyCode: "GBP", Amount: decimal.NewFromFloat(897.01), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{UserId: 15, CurrencyCode: "USD", Amount: decimal.NewFromFloat(87.67), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{UserId: 16, CurrencyCode: "EUR", Amount: decimal.NewFromFloat(312.21), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{UserId: 17, CurrencyCode: "GBP", Amount: decimal.NewFromFloat(0.00), CreatedAt: time.Now(), UpdatedAt: time.Now()},
		}
		db.CreateInBatches(data, 20)
	}

}
