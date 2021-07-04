package repository

import (
	"github.com/matthewagius/wallet-api/entity"
	"gorm.io/gorm"
)

type WalletMySQL struct {
	db *gorm.DB
}

func NewWalletMySQL(db *gorm.DB) *WalletMySQL {
	return &WalletMySQL{
		db: db,
	}
}

func (r *WalletMySQL) Get(id uint) (*entity.Wallet, error) {
	var wallet entity.Wallet
	err := r.db.First(&wallet, id).Error
	if err != nil {

		return nil, err
	}

	return &wallet, nil
}

func (r *WalletMySQL) Update(e *entity.Wallet) error {
	err := r.db.Save(e).Error
	if err != nil {

		return err
	}
	return nil
}
